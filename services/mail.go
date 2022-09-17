package services

import (
	"bytes"
	"fmt"
	"html/template"
	"strconv"

	"github.com/afroluxe/afroluxe-be/config"
	gomail "gopkg.in/mail.v2"
)

type Mail struct {
	From     string
	To       []string
	Subject  string
	Data     interface{}
	Filename string
	Attach   []string
}

var env = config.LoadEnv()

func SendEmail(mail Mail) error {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", mail.From, "AFROLUXE")
	m.SetHeader("To", mail.To...)
	m.SetHeader("Subject", mail.Subject)
	if len(mail.Attach) > 0 {
		for _, filename := range mail.Attach {
			m.Attach(filename)
		}
	}
	parsedHtml := HtmlToString(mail.Filename, mail.Data)
	port, err := strconv.Atoi(env.MailPort)
	if err != nil {
		return err
	}
	m.SetBody("text/html", parsedHtml)
	d := gomail.NewDialer(env.MailHost, port, mail.From, env.MailPassword)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
func HtmlToString(filename string, data interface{}) string {
	t, err := template.ParseFiles(filename)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	buf := new(bytes.Buffer)

	err = t.Execute(buf, data)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	return buf.String()
}
