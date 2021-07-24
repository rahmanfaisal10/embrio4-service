package service

import (
	"crypto/tls"
	"fmt"
	"net"
	"rahmanfaisal10/embrio4-service/config"
	"rahmanfaisal10/embrio4-service/pkg/request"
	"rahmanfaisal10/embrio4-service/pkg/response"
	"regexp"
	"strings"

	"github.com/labstack/gommon/log"
	gomail "gopkg.in/mail.v2"
)

var (
	emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	URL        = "http://localhost:8080/api/change-password"
)

func (s *service) EmailVerificationService(request request.EmailVerificationRequest) *response.BaseResponse {
	//validate format email
	if !isEmailValid(request.Mail) {
		log.Error("not validate email")
		return &response.BaseResponse{
			Success: false,
			Message: fmt.Sprintf("%s is not valid", request.Mail),
		}
	}

	//send to email
	cfg := config.Get()
	message := fmt.Sprintf("please, change password with click this url: \n %s?username_pn=%s", URL, request.UsernamePN)
	m := gomail.NewMessage()

	m.SetHeader("From", cfg.Email)

	// Set E-Mail receivers
	m.SetHeader("To", request.Mail)

	// Set E-Mail subject
	m.SetHeader("Subject", "EMBRIO4 CHANGE PASSWORD MANTRI")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", message)

	// Settings for SMTP server
	d := gomail.NewDialer(cfg.SMTPHost, cfg.SMTPPort, cfg.Email, cfg.MailPassword)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return &response.BaseResponse{
		Success: true,
		Message: "email sent Successfully",
	}
}

func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	if !emailRegex.MatchString(e) {
		return false
	}
	parts := strings.Split(e, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		return false
	}
	return true
}
