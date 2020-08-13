package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Seen        bool               `json:"seen" bson:"seen"`       // notifications.
	Sender      string             `json:"sender" bson:"sender"`   // user.
	Reciver     string             `json:"reciver" bson:"reciver"` // group or user.
	Message     string             `json:"message" bson:"message"`
	DateCreated int64              `json:"dateCreated" bson:"dateCreated"` // Unix time.
}
