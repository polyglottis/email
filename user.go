package email

import (
	"fmt"
	"net/smtp"
)

type User struct {
	Username string
	Password string
	Server   string
	Port     int
}

func (u *User) SendMail(from string, content []byte, to ...string) error {
	auth := smtp.PlainAuth("", u.Username, u.Password, u.Server)
	return smtp.SendMail(fmt.Sprintf("%s:%d", u.Server, u.Port),
		auth, from, to, content)
}
