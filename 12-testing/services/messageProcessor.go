package services

/*

//Not testable

type MessageProcessor struct {
}

var smsService = &SmsService{}

func (m *MessageProcessor) Process(message string) {
	smsService.Send(message)
}
*/

//contract
type MessageService interface {
	Send(message string) bool
}

type MessageProcessor struct {
	msgService MessageService
}

func (m MessageProcessor) Process(message string) {
	m.msgService.Send(message)
	// m.msgService.Send("some random message")
}

func NewMessageProcessor(msgService MessageService) *MessageProcessor {
	return &MessageProcessor{
		msgService: msgService,
	}
}
