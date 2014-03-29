package recvknocking

import (
	"net"
	"sync"
	"time"
)

type Record struct {
	history []time.Time
	fire    func()
	config  recordConfig
	sync.Mutex
}

func NewRecord(handler func(), config recordConfig) Record {
	return Record{
		history: make([]time.Time, 0, int(config.GetCount())),
		fire:    handler,
		config:  config,
	}
}

func (r *Record) Add() {
	r.Lock()
	defer r.Unlock()

	now := time.Now()
	r.history = append(r.history, now)

	dur := r.config.GetDuration()
	for len(r.history) > 1 && now.Sub(r.history[0]) > dur {
		r.history = r.history[1:]
	}

	if len(r.history) < int(r.config.GetCount()) {
		return
	}

	r.fire()
	r.history = r.history[0:0]
}

type Recorder struct {
	records map[string]Record
	config  recorderConfig
	sync.Mutex
}

func NewRecorder(config recorderConfig) Recorder {
	return Recorder{
		records: make(map[string]Record),
		config:  config,
	}
}

func (r *Recorder) Record(addr net.Addr) {
	var ip net.IP
	switch ret := addr.(type) {
	case *net.TCPAddr:
		ip = ret.IP
	case *net.UDPAddr:
		ip = ret.IP
	case *net.IPAddr:
		ip = ret.IP
	default:
		return
	}

	r.Lock()
	defer r.Unlock()

	record, ok := r.records[ip.String()]
	if !ok {
		handler := func() {
			handler := r.config.GetHandler()
			handler(ip)
		}
		record = NewRecord(handler, r.config)
	}

	record.Add()
	r.records[ip.String()] = record
}
