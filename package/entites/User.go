package entites

import "go.mongodb.org/mongo-driver/bson/primitive"

type Users struct {
	ID          primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Username    string             `json:"username" bson:"username"`
	Email       string             `json:"email" bson:"email"`
	PhoneNumber int                `json:"phone_number" bson:"phone_number"`
	Address     string             `json:"address" bson:"address"`
}
