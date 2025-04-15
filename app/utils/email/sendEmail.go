package email

import (
	"context"
	html "html/template"
	"log"
	"os"
	"strconv"
	text "text/template"
	"time"

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

func sendEmail(recipient emailRecipient, subject string, template *template, data any) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	m := mail.NewMsg()
	if err := m.FromFormat("Mazza", os.Getenv("EMAIL_USERNAME")); err != nil {
		log.Default().Printf("failed to set From address: %s", err)
		return err
	}
	if err := m.AddToFormat(recipient.Name, recipient.Address); err != nil {
		log.Default().Printf("failed to set To address: %s", err)
		return err
	}

	m.Subject(subject)

	if err := m.SetBodyTextTemplate(template.Text, data); err != nil {
		log.Default().Printf("failed to set email body: %s", err)
		return err
	}
	if err := m.AddAlternativeHTMLTemplate(template.Html, data); err != nil {
		log.Default().Printf("failed to set email body: %s", err)
		return err
	}

	port, _ := strconv.Atoi(os.Getenv("EMAIL_PORT"))
	c, err := mail.NewClient(os.Getenv("MAIL_HOST"),
		mail.WithPort(port),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(os.Getenv("EMAIL_USERNAME")),
		mail.WithPassword(os.Getenv("EMAIL_PWD")),
		mail.WithSSL(),
		mail.WithSSLPort(true),
		// mail.WithTLSPortPolicy(mail.TLSMandatory),
		// mail.WithTLSConfig(&tls.Config{InsecureSkipVerify: true}),
	)
	if err != nil {
		log.Default().Printf("failed to create email client: %s", err)
		return err
	}
	// fmt.Println(os.Getenv("MAIL_HOST"), port, os.Getenv("EMAIL_USERNAME"), os.Getenv("EMAIL_PWD"), " rec:", recipient.Address)

	if err := c.DialAndSendWithContext(ctx, m); err != nil {
		log.Default().Println("failed to send mail:", err)
		return err
	}

	return nil
}
