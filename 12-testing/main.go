package main

import "testing-app/services"

func main() {
	/*
		msgProcessor := &services.MessageProcessor{}
		msgProcessor.Process("App executed successfully!")
	*/

	//msgProcessor := &services.MessageProcessor{}
	/*
		smsService := &services.SmsService{}
		msgProcessor := services.NewMessageProcessor(smsService)
	*/

	emailService := &services.EmailService{}
	msgProcessor := services.NewMessageProcessor(emailService)
	msgProcessor.Process("App executed successfully!")
}
