package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/sadnamSakib/echo-chat/internal/db"
	"github.com/sadnamSakib/echo-chat/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const roomCollection = "rooms"

func CreateRoom(room *models.Room) error {
	collection := db.MongoDatabase.Collection(roomCollection)
	_, err := collection.InsertOne(context.Background(), room)
	return err
}

func DeleteRoom(room *models.Room) error {
	collection := db.MongoDatabase.Collection(roomCollection)
	_, err := collection.DeleteOne(context.Background(), room)
	return err
}
func GetMessages(roomId string) ([]models.Message, error) {
	fmt.Println("Reached here", roomId)
	collection := db.MongoDatabase.Collection(roomCollection)
	var room models.Room
	err := collection.FindOne(context.Background(), bson.M{"roomid": roomId}).Decode(&room)
	if err == mongo.ErrNoDocuments {
		fmt.Println(err)
		return nil, errors.New("room not found")
	}
	fmt.Println(room)
	var messages []models.Message
	for _, message := range room.Messages {
		messages = append(messages, message)
	}
	return messages, nil

}

func AddUserToRoom(roomId string, email string) error {
	collection := db.MongoDatabase.Collection(roomCollection)
	_, err := collection.UpdateOne(context.Background(), bson.M{"roomId": roomId}, bson.M{"$push": bson.M{"users": email}})
	return err
}

func SendMessage(roomId string, message models.Message) error {
	collection := db.MongoDatabase.Collection(roomCollection)
	_, err := collection.UpdateOne(context.Background(), bson.M{"roomid": roomId}, bson.M{"$push": bson.M{"messages": message}})
	return err
}
