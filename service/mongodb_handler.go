package service

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var CNX = Connection()

func Connection() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://root:Vnpt123@46.137.201.240:27017/cas?authSource=admin")
	// Connect to MongoDB
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	ctx := context.Background()
	client, err := mongo.Connect(ctx, clientOptions)
	//defer client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}
