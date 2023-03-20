package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
    webPort = "80"
    rpcPort = "5001"
    mongoURL = "mongodb://mongo:27017"
    grpcPort = "50001"
)

var client *mongo.Client

type Client struct {
    
}

func main() {
    mongoClient, err := connectToMongo()
    if err != nil {
        log.Panic(err)
    }

    client = mongoClient

    // create a context to disconnect
    ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
    defer cancel()

    // close connection
    defer func ()  {
        if err = client.Disconnect(ctx); err != nil {
            panic(err)
        }
    }()
}

func connectToMongo() (*mongo.Client, error) {
    clientOptions := options.Client().ApplyURI(mongoURL)
    clientOptions.SetAuth(options.Credential{
        Username: "admin",
        Password: "password",
    })

    // mongo connection
    c, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Println("Error connecting:", err)
        return nil, err
    }

    return c, nil
}
