//go:build go1.21
// +build go1.21

package core

import "net"

func setMPTCPForDialer(d *net.Dialer) {
	d.SetMultipathTCP(true)
}

func setMPTCPForListener(lc *net.ListenConfig) {
	lc.SetMultipathTCP(true)
}
