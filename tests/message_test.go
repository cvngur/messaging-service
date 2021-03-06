package tests

import (
	"cvngur/messaging-service/app/services"
	"cvngur/messaging-service/domain"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_should_return_nil_when_user_send_message(t *testing.T) {

	//Arrange
	mockRepo := new(MockMessageRepository)
	mService := services.NewMessageService(mockRepo)
	mockRepo.On("SendMessage").Return(nil)

	var blockedUsers []string
	mockRepo.On("GetBlockedUsers").Return(blockedUsers)

	//Act
	sut := mService.SendMessage("Can", "Ali", "Hello", "01.01.2021")

	//Assert
	mockRepo.AssertExpectations(t)
	assert.Equal(t, nil, sut)
}
func Test_should_return_error_when_blocked_user_send_message(t *testing.T) {

	//Arrange
	mockRepo := new(MockMessageRepository)
	mService := services.NewMessageService(mockRepo)

	blockedUsers := []string{"Can", "Burak"}
	mockRepo.On("GetBlockedUsers").Return(blockedUsers)

	//Act
	sut := mService.SendMessage("Can", "Ali", "Hi", "01.01.2021")

	//Assert
	mockRepo.AssertExpectations(t)
	assert.EqualError(t, sut, "cannot message to user")
}
func Test_should_return_messages(t *testing.T) {

	//Arrange
	mockRepo := new(MockMessageRepository)
	mService := services.NewMessageService(mockRepo)

	messages := []domain.Message{{"Can", "Ali", "Hello", "01.01.2021"},
		{"Ali", "Can", "Hi", "01.01.2021"}}
	mockRepo.On("GetMessages").Return(messages, nil)

	//Act
	sut, err := mService.ViewMessages("Can")

	//Assert
	mockRepo.AssertExpectations(t)
	assert.Equal(t, nil, err)
	assert.NotNil(t, sut)
}
func Test_should_return_error_when_cannot_get_messages(t *testing.T) {

	//Arrange
	mockRepo := new(MockMessageRepository)
	mService := services.NewMessageService(mockRepo)
	var messages []domain.Message
	mockRepo.On("GetMessages").Return(messages, errors.New(""))

	//Act
	sut, err := mService.ViewMessages("Can")

	//Assert
	mockRepo.AssertExpectations(t)
	assert.EqualError(t, err, "cannot view messages")
	assert.Nil(t, sut)
}
func Test_should_return_error_when_cannot_send_message(t *testing.T) {

	//Arrange
	mockRepo := new(MockMessageRepository)
	mService := services.NewMessageService(mockRepo)
	blockedUsers := []string{"Halil", "Burak"}
	mockRepo.On("GetBlockedUsers").Return(blockedUsers)
	mockRepo.On("SendMessage").Return(errors.New(""))

	//Act
	sut := mService.SendMessage("Can", "Ali", "Hi", "04.02.1981")

	//Assert
	mockRepo.AssertExpectations(t)
	assert.EqualError(t, sut, "cannot send message")
}
