//go:build go1.21
// +build go1.21

package core

import (
	"crypto/tls"
	"net"
)

func setMPTCPForDialer(d *net.Dialer) {
	d.SetMultipathTCP(true)
}

func setMPTCPForListener(lc *net.ListenConfig) {
	lc.SetMultipathTCP(true)
}

func isMPTCP(c net.Conn) bool {
	switch tc := c.(type) {
	case *net.TCPConn:
		mp, _ := tc.MultipathTCP()
		return mp
	case *tls.Conn:
		mp, _ := tc.NetConn().(*net.TCPConn).MultipathTCP()
		return mp
	default:
		return false
	}
}
