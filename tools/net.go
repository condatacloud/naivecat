package tools

import (
	"net"
	"time"
)

type INet interface {
	ScanPort(protocol string, hostname string, port string) bool
}

type cnet struct{}

var Net INet = &cnet{}

func (n *cnet) ScanPort(protocol string, hostname string, port string) bool {
	addr := net.JoinHostPort(hostname, port)
	conn, err := net.DialTimeout(protocol, addr, 2*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}
