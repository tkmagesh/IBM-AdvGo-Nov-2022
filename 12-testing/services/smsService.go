package services

import "fmt"

type SmsService struct {
}

func (smsService *SmsService) Send(message string) bool {
	fmt.Printf("Message %q is sent as an sms\n", message)
	return true
}
