package repository

import (
	"context"
	"errors"

	"github.com/sadnamSakib/echo-chat/internal/db"
	"github.com/sadnamSakib/echo-chat/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const userCollection = "users"

// SaveUser saves a new user in the MongoDB
func SaveUser(user *models.User) error {
	collection := db.MongoDatabase.Collection(userCollection)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	_, err = collection.InsertOne(context.Background(), user)
	return err
}

// FindUserByEmail retrieves a user by their email from MongoDB
func FindUserByEmail(email string) (*models.User, error) {
	collection := db.MongoDatabase.Collection(userCollection)
	var user models.User

	err := collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("user not found")
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}

// ComparePasswords compares a hashed password with a plain text password
func ComparePasswords(hashedPassword, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}
