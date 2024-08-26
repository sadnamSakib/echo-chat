package service

import (
	"errors"

	"github.com/sadnamSakib/echo-chat/internal/repository"
	"github.com/sadnamSakib/echo-chat/pkg/models"
)

func SendMessage(message *models.ChatMessage) error {
	// Business logic for sending a message
	if message.Content == "" {
		return errors.New("message content is empty")
	}
	return repository.SaveMessage(message)
}

func ReceiveMessages() ([]models.ChatMessage, error) {
	// Business logic for receiving messages
	return repository.GetMessages()
}
