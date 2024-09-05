package models

type Chat struct {
	ID         string `bson:"_id,omitempty"`
	SenderID   string `bson:"sender_id"`
	ReceiverID string `bson:"receiver_id"`
	QueueName  string `bson:"queue_name"`
}
