package email

import (
	html "html/template"
	"log"
	"os"
	"strconv"
	text "text/template"

	"github.com/wneessen/go-mail"
)

type emailRecipient struct {
	Name    string
	Address string
}

type template struct {
	Html *html.Template
	Text *text.Template
}

func sendEmail(recipient emailRecipient, subject string, template *template, data interface{}) {
	m := mail.NewMsg()
	if err := m.FromFormat("Mazza", "support@mazza.jp"); err != nil {
		log.Default().Printf("failed to set From address: %s", err)
		return
	}
	if err := m.AddToFormat(recipient.Name, recipient.Address); err != nil {
		log.Default().Printf("failed to set To address: %s", err)
		return
	}

	m.Subject(subject)

	if err := m.SetBodyTextTemplate(template.Text, data); err != nil {
		log.Default().Printf("failed to set email body: %s", err)
		return
	}
	if err := m.AddAlternativeHTMLTemplate(template.Html, data); err != nil {
		log.Default().Printf("failed to set email body: %s", err)
		return
	}

	port, _ := strconv.Atoi(os.Getenv("EMAIL_PORT"))
	c, err := mail.NewClient(os.Getenv("MAIL_HOST"), mail.WithPort(port),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(os.Getenv("EMAIL_USERNAME")),
		mail.WithPassword(os.Getenv("EMAIL_PWD")),
		mail.WithTLSPortPolicy(mail.TLSMandatory),
	)
	if err != nil {
		log.Default().Printf("failed to create email client: %s", err)
		return
	}
	// defer c.Close()

	if err := c.DialAndSend(m); err != nil {
		log.Default().Printf("failed to send mail: %s", err)
	}
}
