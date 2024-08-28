package models

type AuthToken struct {
	ID        string `bson:"_id,omitempty"`
	Token     string `bson:"token"`
	UserID    string `bson:"user_id"`
	ExpiresAt int64  `bson:"expires_at"`
}
