package shadowio

import "net"

type WaitRead interface {
	WaitRead() (data []byte, put func(), err error)
}

type WaitReadFrom interface {
	WaitReadFrom() (data []byte, put func(), addr net.Addr, err error)
}
