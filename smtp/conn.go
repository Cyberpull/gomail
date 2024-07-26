package smtp

import (
	"bufio"
	"net"
	"sync"
)

type Conn struct {
	wm     sync.Mutex
	rm     sync.Mutex
	conn   net.Conn
	reader *bufio.Reader
}

func (x *Conn) Read(b []byte) (n int, err error) {
	x.rm.Lock()

	defer x.rm.Unlock()

	return x.conn.Read(b)
}

func (x *Conn) Write(b []byte) (n int, err error) {
	x.wm.Lock()

	defer x.wm.Unlock()

	return x.conn.Write(b)
}

func (x *Conn) Close() (err error) {
	if x.conn != nil {
		err = x.conn.Close()
	}

	return
}

// ======================

func newConn(conn net.Conn) *Conn {
	value := &Conn{conn: conn}
	value.reader = bufio.NewReader(value)
	return value
}
