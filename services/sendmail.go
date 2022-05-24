package services

import (
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailObject struct {
	To      string
	Body    string
	Subject string
}

func SendMail(subject string, body string, to string, name string, html string) bool {
	from := mail.NewEmail("Open it", os.Getenv(""))
	_to := mail.NewEmail(name, to)
	TextContent := body
	htmlContent := html
	message := mail.NewSingleEmail(from, subject, _to, TextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv(""))
	_, err := client.Send(message)
	if err != nil {
		return false
	} else {
		return true
	}

}
