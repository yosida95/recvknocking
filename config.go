package recvknocking

import (
	"net"
	"time"
)

type Factory func() (net.Listener, error)

type Handler func(net.IP)

type Config struct {
	Count    uint
	Duration time.Duration
	Factory  Factory
	Handler  Handler
}

func (c Config) GetCount() uint {
	return c.Count
}

func (c Config) GetDuration() time.Duration {
	return c.Duration
}

func (c Config) GetFactory() Factory {
	return c.Factory
}

func (c Config) GetHandler() Handler {
	return c.Handler
}

type receiverConfig interface {
	GetFactory() Factory
}

type recorderConfig interface {
	GetHandler() Handler
	GetCount() uint
	GetDuration() time.Duration
}

type recordConfig interface {
	GetCount() uint
	GetDuration() time.Duration
}
