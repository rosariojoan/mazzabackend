package email

import (
	"fmt"
	html "html/template"
	"log"
	text "text/template"
)

func SendPasswordResetToken(name string, emailAddress string, token string) {
	subject := "Código de Reinício de Palavra-Passe"
	data := struct {
		Name  string
		Token string
	}{
		Name:  name,
		Token: token,
	}

	htmlTemplate, err := html.ParseFiles("./app/utils/email/templates/password_reset.html")
	if err != nil {
		log.Default().Printf("1. failed to load email template: %s", err)
		return
	}

	textTemplate, err := text.ParseFiles("./app/utils/email/templates/password_reset.html")
	if err != nil {
		log.Default().Printf("2. failed to load email template: %s", err)
		return
	}

	err = sendEmail(
		emailRecipient{Name: name, Address: emailAddress},
		subject,
		&template{Html: htmlTemplate, Text: textTemplate},
		data,
	)

	fmt.Println("err:", err)
}
