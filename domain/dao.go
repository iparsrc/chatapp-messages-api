package domain

import (
	"context"
	"time"

	"github.com/parsaakbari1209/ChatApp-messages-api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Create(msg *Message) (*Message, *utils.RestErr) {
	msgC := db.Collection("messages")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, err := msgC.InsertOne(ctx, bson.M{
		"seen":        msg.Seen,
		"sender":      msg.Sender,
		"reciver":     msg.Reciver,
		"message":     msg.Message,
		"dateCreated": msg.DateCreated,
	})
	if err != nil {
		return nil, utils.InternalServerErr("cant' operate create functionality.")
	}
	if res.InsertedID == nil {
		return nil, utils.InternalServerErr("can't operate create functionality.")
	}
	msg.ID = res.InsertedID.(primitive.ObjectID)
	return msg, nil
}
