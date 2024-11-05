package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Note struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title     string             `bson:"title" json:"title"`
	Content   string             `bson:"content" json:"content"`
	Link      string             `bson:"link" json:"link"`
	ExpiredAt primitive.DateTime `bson:"expired_at" json:"expired_at"`
}
