package mongodb

import (
	"context"
	"github.com/xu756/imlogic/common/types"
	"github.com/xu756/imlogic/kitex_gen/im"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

var (
	database   = "imlogic"
	collection = "msg"
)

type MsgApi interface {
	CreateMsg(ctx context.Context, msg *im.Message) (err error)
	QueryMsg(ctx context.Context, from, to string) (res *[]im.Message, err error)
}

func (m *customClient) CreateMsg(ctx context.Context, msg *im.Message) (err error) {
	err = m.client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}
	_, err = m.client.Database(database).Collection(collection).InsertOne(ctx, msg)
	return
}

func (m *customClient) QueryMsg(ctx context.Context, from, to string) (res *[]im.Message, err error) {

	err = m.client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	var result types.Message
	err = m.client.Database(database).
		Collection(collection).
		FindOne(ctx, bson.D{{"from", from}, {"to", to}}).Decode(&result)
	if err != nil {
		return nil, err
	}
	log.Print(result)
	return
}
