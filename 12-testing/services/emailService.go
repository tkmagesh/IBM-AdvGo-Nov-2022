package services

import "fmt"

type EmailService struct {
}

func (emailService *EmailService) Send(message string) bool {
	fmt.Printf("Message %q is sent as an e-mail\n", message)
	return true
}
