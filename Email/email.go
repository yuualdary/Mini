package Email

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"time"
)

func SendEmail(name string, code int, valid time.Time, email string) bool{

  // Sender data.
  from := "magangku21@gmail.com"
  fmt.Println(from)
  password :=  "2021magangku"

  // Receiver email address.
  to := []string{
    email,
  }

  // smtp server configuration.
  smtpHost := "smtp.gmail.com"
  smtpPort := "587"

  // Message.
 // message := []byte("This is a test email message.")
  
  // Authentication.
  auth := smtp.PlainAuth("", from, password, smtpHost)
  

  t, _ := template.ParseFiles("emailtemplate.html")


  var body bytes.Buffer

  mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
  body.Write([]byte(fmt.Sprintf("Subject:Email Verification \n%s\n\n", mimeHeaders)))

  t.Execute(&body, struct {
    Name    string
    Code int
	Valid time.Time
	Email string
  }{
    Name:    name,
    Code: code,
	Valid: valid,
	Email: email,
  })
  //fixing kalau dia gagal kirim email, gabisa berkali" register blm berjalan

  // Sending email.
  err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to,  body.Bytes())
  if err != nil {
    fmt.Println(err)
    return false
  }
  return true
}