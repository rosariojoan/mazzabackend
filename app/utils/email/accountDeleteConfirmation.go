package email

import (
	"fmt"
	html "html/template"
	"log"
	"os"
	text "text/template"
	"time"
)

func SendAccountDeleteConfirmation(name string, emailAddress string, token string, validity string) {
	subject := "Eliminação da Conta"
	data := struct {
		Link        string
		Year        int
		CompanyName string
		Validity    string
	}{
		Link:        fmt.Sprintf("%s/account/unsubscribe?token=%s", os.Getenv("WEB_HOME_PAGE"), token),
		Year:        time.Now().Year(),
		CompanyName: os.Getenv("COMPANY_NAME"),
		Validity:    validity,
	}

	var filePath string
	if os.Getenv("DEBUG") == "true" {
		filePath = "./app/utils/email/templates/account_delete_request.html"
	} else {
		filePath = "./app/build/utils/email/templates/account_delete_request.html"
	}
	htmlTemplate, err := html.ParseFiles(filePath)
	if err != nil {
		log.Default().Printf("1. failed to load email template: %s", err)
		return
	}

	textTemplate, err := text.ParseFiles(filePath)
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
