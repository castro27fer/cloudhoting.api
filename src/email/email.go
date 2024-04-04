package email

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"net/mail"
	"net/smtp"
	"os"

	language "github.com/ebarquero85/link-backend/src/translations"
)

type Dest struct {
	Name string
	Code string
}

func SendActivationEmail(to mail.Address, code string) (err error) {

	translator := language.Get_translator()
	subject, err := translator.T("user_register_subject")
	if err != nil {
		subject = "Cloud hosting account activation"
	}

	if err = SendEmail(to, subject, code); err != nil {
		return err
	}

	return nil
}

func SendEmail(to mail.Address, subject string, Code string) (err error) {

	var (
		email          = os.Getenv("EMAIL")
		email_name     = os.Getenv("EMAIL_NAME")
		email_password = os.Getenv("EMAIL_PASSWORD")
		host           = os.Getenv("EMAIL_HOST")
		port           = os.Getenv("EMAIL_HOST_PORT")
	)

	from := mail.Address{Name: email_name, Address: email}

	dest := Dest{Name: to.Name, Code: Code}

	message := ""

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject
	headers["Content-Type"] = `text/html; charset="UTF-8"`

	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	t, err := template.ParseFiles("./src/email/template.html")
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, dest)
	if err != nil {
		return err
	}

	message += buf.String()

	auth := smtp.PlainAuth("", email, email_password, host)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", host+":"+port, tlsConfig)
	if err != nil {
		return err
	}

	client, err := smtp.NewClient(conn, host)
	if err != nil {
		return err
	}

	err = client.Auth(auth)
	if err != nil {
		return err
	}

	err = client.Mail(from.Address)
	if err != nil {
		return err
	}

	err = client.Rcpt(to.Address)
	if err != nil {
		return err
	}

	w, err := client.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	err = client.Quit()
	if err != nil {
		return err
	}

	return nil
}
