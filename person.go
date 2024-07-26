package gomail

import (
	"fmt"
	"strings"
)

type Person struct {
	Name  string
	Email string
}

func (x *Person) toEmailWithName() string {
	value := x.toEmail()

	x.Name = strings.TrimSpace(x.Name)

	if x.Name != "" {
		value = fmt.Sprintf(`"%s" %s`, x.Name, value)
	}

	return value
}

func (x *Person) toEmail() string {
	if x.Email == "" {
		return x.Email
	}

	return fmt.Sprintf("<%s>", x.Email)
}

func (x *Person) clear() {
	x.Name = ""
	x.Email = ""
}
