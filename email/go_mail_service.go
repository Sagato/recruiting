package email

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"gopkg.in/gomail.v2"
)

type GoMailService struct {
	Smtp       string
	Port       int
	User       string
	From       string
	Subject    string
	Password   string
	Recipients []string
}

func NewGoMailService() *GoMailService {
	smtp := os.Getenv("SMTP")
	user := os.Getenv("EMAIL_USER")
	password := os.Getenv("EMAIL_PASSWORD")
	from := os.Getenv("FROM")
	subject := "Contact Us Form"
	recipients := strings.Split(os.Getenv("RECIPIENTS"), ",")
	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		panic("Port is not an integer value")
	}

	return &GoMailService{
		Smtp:       smtp,
		Port:       port,
		User:       user,
		From:       from,
		Subject:    subject,
		Password:   password,
		Recipients: recipients,
	}
}

func (gm *GoMailService) SendMail(body string) error {
	d := gomail.NewDialer(gm.Smtp, gm.Port, gm.User, gm.Password)
	s, err := d.Dial()
	if err != nil {
		return err
	}

	fmt.Println("Mail Service", gm.User)
	m := gomail.NewMessage()
	for _, recipient := range gm.Recipients {
		m.SetHeader("From", gm.User)
		m.SetHeader("Subject", gm.Subject)
		m.SetBody("text/html", body)
		m.SetAddressHeader("To", recipient, recipient)

		if err := gomail.Send(s, m); err != nil {
			log.Printf("Could not send email to %q: %v", recipient, err)
		}
		m.Reset()
	}
	return nil
}
