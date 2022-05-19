package utils

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
)

func BaseSend(location string, detail string) {
	mailHeader := map[string][]string{
		"From":    {"applescriptserver@163.com"},
		"To":      {"applescriptserver@163.com"},
		"Subject": {"标题"},
	}

	m := gomail.NewMessage()
	m.SetHeaders(mailHeader)
	m.SetBody("text/html", fmt.Sprintf(mapUrl, location, location, mkey)+"<br>"+detail+"<br>"+location)

	d := gomail.NewDialer("smtp.163.com", 465, "applescriptserver@163.com", "GUDWSQQCTRKVGACA")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error" + err.Error())
	}
}
