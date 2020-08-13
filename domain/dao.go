package domain

import (
	"context"
	"time"

	"github.com/parsaakbari1209/ChatApp-messages-api/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func Create(msg *Message) (*Message, *utils.RestErr) {
	msgC := db.Collection("messages")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, err := msgC.InsertOne(ctx, bson.M{
		"_id":         msg.ID,
		"seen":        msg.Seen,
		"sender":      msg.Sender,
		"reciver":     msg.Reciver,
		"message":     msg.Message,
		"dateCreated": msg.DateCreated,
	})
	if err != nil {
		return nil, utils.InternalServerErr("cant' operate create functionality.")
	}
	if res.InsertedID != msg.ID {
		return nil, utils.InternalServerErr("can't operate create functionality.")
	}
	return msg, nil
}

func Delete(id string) *utils.RestErr {
	msgC := db.Collection("messages")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, err := msgC.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return utils.InternalServerErr("can't operate delete functionality.")
	}
	if res.DeletedCount == 0 {
		return utils.NotFound("message not found.")
	}
	return nil
}
