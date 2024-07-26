package smtp

import "sync"

type Attachment interface {
	read() (b []byte, err error)
}

type Content struct {
	mutex      sync.Mutex
	attachment []Attachment
	subject    string
	html       string
	text       string
}

func (x *Content) Attachment(attachments ...Attachment) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	x.subject = subject
}

func (x *Content) Subject(subject string) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	x.subject = subject
}

func (x *Content) HTML(html string) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	x.html = html
}

func (x *Content) Text(text string) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	x.text = text
}

func (x *Content) validate() (err error) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	return
}
