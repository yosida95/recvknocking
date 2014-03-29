package recvknocking

import (
	"net"
	"time"
)

type factory func() (net.Listener, error)

type handler func(net.IP)

type Config struct {
	Count    uint
	Duration time.Duration
	Factory  factory
	Handler  handler
}

func (c Config) GetCount() uint {
	return c.Count
}

func (c Config) GetDuration() time.Duration {
	return c.Duration
}

func (c Config) GetFactory() factory {
	return c.Factory
}

func (c Config) GetHandler() handler {
	return c.Handler
}

type receiverConfig interface {
	GetFactory() factory
}

type recorderConfig interface {
	GetHandler() handler
	GetCount() uint
	GetDuration() time.Duration
}

type recordConfig interface {
	GetCount() uint
	GetDuration() time.Duration
}
