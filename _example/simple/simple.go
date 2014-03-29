package main

import (
	"log"
	"net"
	"time"

	"github.com/yosida95/recvknocking"
)

func main() {
	config := recvknocking.Config{
		Count:    3,
		Duration: 1 * time.Second,
		Factory: func() (net.Listener, error) {
			return net.Listen("tcp4", ":22222")
		},
		Handler: func(ip net.IP) {
			log.Printf("fire for %s\n", ip.String())
		},
	}

	r := recvknocking.NewReceiver(config)
	r.Run()
}
