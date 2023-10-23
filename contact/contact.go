package contact

import (
	"bytes"
	"fmt"
	"net/http"
	"text/template"

	"github.com/Cheveo/recruiting/email"
	"github.com/Cheveo/recruiting/errors"
	"github.com/gin-gonic/gin"
)

type Message struct {
	Firstname string
	Lastname  string
	Email     string
	Message   string
}

type ContactUsHandler struct {
	MailService email.EmailService
}

func NewContactUsHandler(mailService email.EmailService) *ContactUsHandler {
	return &ContactUsHandler{
		MailService: mailService,
	}
}

func (h *ContactUsHandler) handlePostContact(c *gin.Context) {
	var message Message
	err := c.BindJSON(&message)
	if err != nil {
		httperror := errors.NewHttpError("couldnt serialize data", http.StatusBadRequest)
		c.Error(httperror)
		return
	}

	t, err := template.ParseFiles("templates/contact_us.html")
	if err != nil {
		panic(err.Error())
	}

	var tpl bytes.Buffer
	t.Execute(&tpl, message)
	err = h.MailService.SendMail(tpl.String())
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (h *ContactUsHandler) SetupRouter(r *gin.Engine) {
	r.POST("/contact-us", h.handlePostContact)
}
