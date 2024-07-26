package smtp

import (
	"sync"

	"cyberpull.com/gomail"
)

type Envelope struct {
	mutex   sync.Mutex
	From    gomail.Person
	ReplyTo gomail.Person
	To      []gomail.Person
	Cc      []gomail.Person
	Bcc     []gomail.Person
}

func (x *Envelope) AddTo(persons ...gomail.Person) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	x.initialize()

	for _, person := range persons {
		if person.Email == "" {
			continue
		}

		x.To = append(x.To, person)
	}
}

func (x *Envelope) AddCc(persons ...gomail.Person) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	x.initialize()

	for _, person := range persons {
		if person.Email == "" {
			continue
		}

		x.Cc = append(x.Cc, person)
	}
}

func (x *Envelope) AddBcc(persons ...gomail.Person) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	x.initialize()

	for _, person := range persons {
		if person.Email == "" {
			continue
		}

		x.Bcc = append(x.Bcc, person)
	}
}

func (x *Envelope) validate() (err error) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	return
}

func (x *Envelope) initialize() {
	if x.To == nil {
		x.To = make([]gomail.Person, 0)
	}

	if x.Cc == nil {
		x.Cc = make([]gomail.Person, 0)
	}

	if x.Bcc == nil {
		x.Bcc = make([]gomail.Person, 0)
	}
}
