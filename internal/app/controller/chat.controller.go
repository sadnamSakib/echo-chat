package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sadnamSakib/echo-chat/internal/repository"
	"github.com/sadnamSakib/echo-chat/pkg/models"

	"github.com/labstack/echo/v4"
)

func NewChatRoom(c echo.Context) error {
	// Parse the request to get room details
	room := new(models.Room)
	room.RoomId = uuid.New().String()
	room.Messages = make([]models.Message, 0)
	room.Users = make([]string, 0)
	if err := c.Bind(room); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	// Create the room using repository function
	err := repository.CreateRoom(room)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create room"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Room created successfully", "roomId": room.RoomId})
}

// ReceiveMessages retrieves messages from a specific room
func ReceiveMessages(c echo.Context) error {
	roomId := c.Param("roomId")
	messages, err := repository.GetMessages(roomId)
	fmt.Println(roomId)
	if err != nil {

		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to retrieve room messages"})
	}

	return c.JSON(http.StatusOK, messages)
}

// AddUserToRoom adds a user to a chat room
func AddUserToRoom(c echo.Context) error {
	roomId := c.Param("roomId")
	email := c.FormValue("email")

	if err := repository.AddUserToRoom(roomId, email); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to add user to room"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "User added to room successfully"})
}

// SendMessage handles sending a message to a specific room
func SendMessage(c echo.Context) error {
	roomId := c.Param("roomId")
	message := new(models.Message)
	message.Time = time.Now().Format("2006-01-02 15:04:05")

	if err := c.Bind(message); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	// Append the message to the room's message list using repository function
	err := repository.SendMessage(roomId, *message)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to send message"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Message sent successfully"})
}
