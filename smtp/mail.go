package smtp

type Mail struct {
	Envelope Envelope
	Content  Content
}

func (x *Mail) send(smtp *SMTP) (err error) {
	if err = x.Envelope.validate(); err != nil {
		return
	}

	if err = x.Content.validate(); err != nil {
		return
	}

	defer func() {
		if err != nil {
			smtp.reset()
		}
	}()

	return
}
