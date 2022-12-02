package services

import (
	"testing"
	mocks "testing-app/mocks/services"
)

/*
// custom mock implementation
type MockMessageService struct {
	receivedMessage string
	isCalled        bool
	expectedMessage string
}

func (m *MockMessageService) Send(message string) bool {
	m.isCalled = true
	m.receivedMessage = message
	return true
}

func Test_MessageProcesser_Sends_Message_To_MessageService(t *testing.T) {
	// arrange
	testMessage := "test message"
	mockMessageService := &MockMessageService{
		expectedMessage: testMessage,
	}

	sut := NewMessageProcessor(mockMessageService)

	// act
	sut.Process(testMessage)

	// assert
	if !mockMessageService.isCalled {
		t.Error("Message processor did not send the message")
	}

	if mockMessageService.receivedMessage != testMessage {
		t.Errorf("Expected message : %q to be sent to the message service, instead sent %q\n", testMessage, mockMessageService.receivedMessage)
	}
}
*/

//using "github.com/stretchr/testify/mock" package
func Test_MessageProcesser_Sends_Message_To_MessageService(t *testing.T) {
	// arrange
	testMessage := "test message"
	mockMessageService := &mocks.MessageService{}

	//configure the mock
	mockMessageService.On("Send", testMessage).Return(true)
	sut := NewMessageProcessor(mockMessageService)

	// act
	sut.Process(testMessage)

	// assert
	mockMessageService.AssertExpectations(t)
}
