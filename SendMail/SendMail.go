package sendmail

import (
	"fmt"
	"net/smtp"
	"crypto/tls"
	"net"
	"net/mail"
	"strings"
	"encoding/base64"
	//"mime/multipart"
)

func SendMail(smtpServer string, auth smtp.Auth, from mail.Address, to mail.Address, title string, body string, jpg map[string][]byte) error {
	
	var err error
	smtpHost, _, _ := net.SplitHostPort(smtpServer)

	header := make(map[string]string)
	
	header["Message-ID"] = "PL-001" //TODO
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = encodeRFC2047(title)
	header["MIME-Version"] = "1.0"
	header["Content-Transfer-Encoding"] = "base64"
	
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "Content-Type: Multipart/Mixed; boundary=b01; type=text/html; charset=\"utf-8\"" + "\r\n" + "\r\n"
	
	message += "--b01" + "\r\n" +
		"Content-Type: text/html; charset=\"utf-8\"" + "\n" +
		"Content-Transfer-Encoding: base64" + "\n" + "\r\n" +
		base64.StdEncoding.EncodeToString([]byte(body)) + "\r\n"
	
	for jpgName, jpgBody := range jpg {
		message += "--b01" + "\n" +
			"Content-Type: image/gif" + "\n" +
			"Content-ID: "+ jpgName + "\n" +
			"Content-Transfer-Encoding: base64" + "\n" + "\r\n" +
			base64.StdEncoding.EncodeToString(jpgBody) + "\r\n"
	}
	message += "\r\n" + "--b01--" + "\r\n"
	
	//log.Println(message)
	
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpHost,
	}
	conn, err := tls.Dial("tcp", smtpServer, tlsConfig)
	if err != nil {
		return err
		//log.Fatal(err)
	}
	c, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		return err
		//log.Fatal(err)
	}
	if err := c.Auth(auth); err != nil {
		return err
		//log.Panicln(err)
	}
	
	if err = c.Mail(from.Address); err != nil {
		return err
		//log.Panicln(err)
	}
	if err = c.Rcpt(to.Address); err != nil {
		return err
		//log.Panicln(err)
	}
	if w, err := c.Data(); err != nil {
		return err
		//log.Panicln(err)
	} else {
		defer w.Close()
		if _, err = w.Write([]byte(message)); err != nil {
			return err
		}
		w.Close()
	}
	
	err = c.Quit()
	return err
}

func encodeRFC2047(String string) string {
	addr := mail.Address{String, ""}
	return strings.Trim(addr.String(), " <@>")
}
