package main

import (
	"context"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func mongoConnect() (database *mongo.Client, Context context.Context) {
	USER := os.Getenv("USERR")
	PASS := os.Getenv("PASS")
	PROTOCOL := os.Getenv("PROTOCOLMONGO")
	uri := "mongodb://" + USER + ":" + PASS + "@" + PROTOCOL

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	c, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	defer cancel()
	if err != nil {
		panic(err.Error())
	}
	return c, ctx
}

func envLoad() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
}
