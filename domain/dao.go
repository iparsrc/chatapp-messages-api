package domain

import (
	"context"
	"time"

	"github.com/parsaakbari1209/ChatApp-messages-api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func RetriveTwnety(reciver string, skipCof int64, sender *string) (*[]Message, *utils.RestErr) {
	msgC := db.Collection("messages")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	// Setting options to read the last 40 inserted docs of the related chat.
	ops := options.Find()
	ops.SetSkip(40 * skipCof)
	ops.SetLimit(40)
	ops.SetSort(bson.D{{Key: "dateCreated", Value: -1}})
	// Filter based on a group chat or direct chat.
	filter := bson.M{
		"reciver": reciver,
	}
	if sender != nil {
		filter = bson.M{
			"reciver": reciver,
			"sender":  &sender,
		}
	}
	// Retrive docs.
	cur, err := msgC.Find(ctx, filter, ops)
	if err != nil {
		return nil, utils.InternalServerErr("can't operate find functionality.")
	}
	// Write all retrived docs to a slice of Messages.
	var messages []Message
	ctxx, cancelx := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelx()
	if err := cur.All(ctxx, &messages); err != nil {
		return nil, utils.InternalServerErr(err.Error())
	}
	// Retrun messages and no err.
	return &messages, nil
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
