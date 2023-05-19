package shadowio

import (
	"io"
	"net"
)

type WaitRead interface {
	WaitRead() (data []byte, put func(), err error)
}

type WaitReadReader interface {
	io.Reader
	WaitRead
}

type WaitReadFrom interface {
	WaitReadFrom() (data []byte, put func(), addr net.Addr, err error)
}
