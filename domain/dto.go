package domain

type Message struct {
	ID          string `json:"id" bson:"_id"`          // Sender + DateCreated.
	Seen        bool   `json:"seen" bson:"seen"`       // notifications.
	Sender      string `json:"sender" bson:"sender"`   // user.
	Reciver     string `json:"reciver" bson:"reciver"` // group or user.
	Message     string `json:"message" bson:"message"`
	DateCreated int64  `json:"dateCreated" bson:"dateCreated"` // Unix time.
}
