package smtp

import (
	"sync"

	"cyberpull.com/gokit/errors"
)

type SMTP struct {
	mutex   sync.Mutex
	conn    *Conn
	address string
}

func (x *SMTP) Open(address string) (err error) {
	x.mutex.Lock()

	defer func() {
		x.mutex.Unlock()

		if err != nil {
			x.Close()
		}
	}()

	x.address = address

	return
}

func (x *SMTP) Close() (err error) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	if x.conn != nil {
		x.conn.Close()
		x.conn = nil
	}

	x.address = ""

	return
}

func (x *SMTP) Send(mails ...*Mail) (err error) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	if x.conn == nil {
		err = errors.New("Connection not open")
		return
	}

	for _, mail := range mails {
		if mail == nil {
			err = errors.New("Invalid mail object")
			break
		}

		err = mail.send(x)

		if err != nil {
			break
		}
	}

	return
}

func (x *SMTP) reset() {
	x.mutex.Lock()

	defer x.mutex.Unlock()
}
