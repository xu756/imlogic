package mongodb

import (
	"context"
	"github.com/xu756/imlogic/common/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
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
