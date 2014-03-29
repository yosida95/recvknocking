package recvknocking

import (
	"net"
	"time"
)

type Receiver struct {
	config   receiverConfig
	recorder Recorder
}

func NewReceiver(config Config) *Receiver {
	return &Receiver{
		config:   config,
		recorder: NewRecorder(config),
	}
}

func (r *Receiver) communicate(conn net.Conn) {
	defer conn.Close()

	r.recorder.Record(conn.RemoteAddr())
}

func (r *Receiver) next(err error) bool {
	const (
		max     = 1 * time.Second
		halfmax = max / 2
	)
	var sleepdur time.Duration

	if err != nil {
		if nerr, ok := err.(net.Error); !ok || !nerr.Temporary() {
			return false
		}

		if sleepdur == 0 {
			sleepdur = 5 * time.Millisecond
		} else if sleepdur > halfmax {
			sleepdur = max
		} else {
			sleepdur *= 2
		}
		time.Sleep(sleepdur)
	}

	return true
}

func (r *Receiver) serve(l net.Listener) (err error) {
	defer l.Close()

	for {
		conn, err := l.Accept()
		if !r.next(err) {
			return err
		}

		go r.communicate(conn)
	}
}

func (r *Receiver) Run() (err error) {
	l, err := r.config.GetFactory()()
	if err != nil {
		return
	}

	return r.serve(l)
}
