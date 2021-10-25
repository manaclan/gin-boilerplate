package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/manaclan/gin-boilerplate/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Services struct {
	Client *mongo.Client
}

func (services Services) RegisterUser(username string, password string) (code int, message string) {
	fmt.Println("Service: RegisterUser")
	var users []models.User
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	currentUsersCollection := services.Client.Database("users").Collection("current-users")
	defaultUsersCollection := services.Client.Database("users").Collection("default-users")

	cursor, err := defaultUsersCollection.Find(
		context.Background(), bson.M{"username": username})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	if err = cursor.All(ctx, &users); err != nil {
		panic(err)
	}
	if len(users) > 0 {
		return 404, "User existed!"
	}
	newUser := models.User{Username: username, Password: password}
	newUser.Password, err = models.Hash(newUser.Password)
	if err != nil {
		panic(err)
	}
	_, err = defaultUsersCollection.InsertOne(ctx, newUser)
	if err != nil {
		panic(err)
	}
	_, err = currentUsersCollection.InsertOne(ctx, newUser)
	if err != nil {
		panic(err)
	}
	return 200, "Succeed!"
}
func (services Services) ValidateUser(username string, password string) bool {
	fmt.Println("Service: ValidateUser")
	var users []models.User
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	currentUsersCollection := services.Client.Database("users").Collection("current-users")

	cursor, err := currentUsersCollection.Find(
		context.Background(), bson.M{"username": username})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	if err = cursor.All(ctx, &users); err != nil {
		panic(err)
	}
	if len(users) == 0 {
		return false
	}
	err = models.CheckPasswordHash(users[0].Password, password)
	return err == nil
}
