package services

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"net/mail"
	"net/smtp"
	"os"
)

type FormatEmail struct {
	Name         string
	TemplatePath string
	Description  string
	Fields       []string
}

type Dest struct {
	Name string
	Code string
}

type Email struct {
	Email      string
	Email_Name string
	Password   string
	Host_SMTP  string
	Port       string
}

var EmailService Email = Email{}

func (email *Email) Start() {

	email.Email = os.Getenv("EMAIL")
	email.Email_Name = os.Getenv("EMAIL_NAME")
	email.Password = os.Getenv("EMAIL_PASSWORD")
	email.Host_SMTP = os.Getenv("EMAIL_HOST")
	email.Port = os.Getenv("EMAIL_HOST_PORT")

}

func (email *Email) SendEmail(to mail.Address, subject string, Code string) (err error) {

	from := mail.Address{Name: email.Email_Name, Address: email.Email}

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

	t, err := template.ParseFiles("./src/services/email/template.html")
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, dest)
	if err != nil {
		return err
	}

	message += buf.String()

	auth := smtp.PlainAuth("", email.Email, email.Password, email.Host_SMTP)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         email.Host_SMTP,
	}

	conn, err := tls.Dial("tcp", email.Host_SMTP+":"+email.Port, tlsConfig)
	if err != nil {
		return err
	}

	client, err := smtp.NewClient(conn, email.Host_SMTP)
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
