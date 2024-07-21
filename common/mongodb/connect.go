package mongodb

import (
	"context"
	"imlogic/common/config"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type customClient struct {
	client *mongo.Client
}

func NewClient() Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.RunData.MongodbUrl))
	if err != nil {
		log.Panic(err)
	}
	return &customClient{
		client: client,
	}
}
