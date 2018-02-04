package sendmail

import (
	"testing"
	"net/smtp"
	"net"
	"net/mail"
	"io/ioutil"
	"log"
)

func TestSendMail(T *testing.T) {
	
	smtpServer := "smtp.yandex.ru:465"
	smtpHost, _, _ := net.SplitHostPort(smtpServer)
	
	smtpAuth := smtp.PlainAuth(
		"",
		"username@yandex.ru",
		"password***",
		smtpHost,
	)
	
	from := mail.Address{"UserName", "username@yandex.ru"}
	to := mail.Address{"UserName", "username@yandex.ru"}
	
	jpg := make(map[string][]byte)
	
	var err error
	if jpg["test_jpg"], err = ioutil.ReadFile("test_jpg.jpg"); err != nil {
		log.Panic(err)
	}
	
	if jpg["test_gif"], err = ioutil.ReadFile("test_gif.gif"); err != nil {
		log.Panic(err)
	}
	
	if jpg["test_png"], err = ioutil.ReadFile("test_png.png"); err != nil {
		log.Panic(err)
	}
	
	SendMail(
		smtpServer,
		smtpAuth,
		from,
		to,
		"Title - Тема письма",
		`Body <br><h1>Тело сообщения</h1>
		<br><IMG SRC="cid:test_jpg" BORDER=0 HEIGHT=100 WIDTH=100 ALT="image - картинка 1">
		<br><IMG SRC="cid:test_gif" BORDER=0 HEIGHT=100 WIDTH=100 ALT="image - картинка 2">
		<br><IMG SRC="cid:test_png" BORDER=0 HEIGHT=100 WIDTH=100 ALT="image - картинка 3">
		`,
		jpg,
	)
	
}
