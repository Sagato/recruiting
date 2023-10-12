package email

type EmailService interface {
	SendMail(body string) error
}
