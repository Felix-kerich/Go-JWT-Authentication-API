package services

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

type EmailService struct {
	host     string
	port     string
	username string
	password string
	from     string
}


func NewEmailService() *EmailService {
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")

	// Debug: Print values to confirm they are loaded
	log.Printf("📧 SMTP Config: Host=%s, Port=%s, Username=%s\n", host, port, username)

	
	if host == "" || port == "" || username == "" || password == "" {
		log.Fatal("❌ SMTP Configuration is missing! Check your .env file.")
	}

	return &EmailService{
		host:     host,
		port:     port,
		username: username,
		password: password,
		from:     username, // Sender email
	}
}


func (s *EmailService) SendVerificationEmail(to, token string) error {
	verificationLink := fmt.Sprintf("%s/auth/verify-email?token=%s", os.Getenv("APP_URL"), token)

	subject := "Verify Your Email"
	body := fmt.Sprintf(`
		<html>
			<body>
				<h2>Welcome to Our Platform!</h2>
				<p>Please click the link below to verify your email:</p>
				<a href="%s">Verify Email</a>
				<p>If you didn't create an account, please ignore this email.</p>
			</body>
		</html>
	`, verificationLink)

	return s.sendEmail(to, subject, body)
}

func (s *EmailService) sendEmail(to, subject, body string) error {
	msg := fmt.Sprintf("To: %s\r\n"+
		"Subject: %s\r\n"+
		"MIME-Version: 1.0\r\n"+
		"Content-Type: text/html; charset=UTF-8\r\n"+
		"\r\n"+
		"%s\r\n", to, subject, body)

	auth := smtp.PlainAuth("", s.username, s.password, s.host)
	addr := fmt.Sprintf("%s:%s", s.host, s.port)

	return smtp.SendMail(addr, auth, s.from, []string{to}, []byte(msg))
}
