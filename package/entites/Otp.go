package entites

import "go.mongodb.org/mongo-driver/bson/primitive"

type Otp struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	OtpCode int                `json:"otp_code" bson:"otp_code"`
	UserID  primitive.ObjectID `json:"user_id" bson:"user_id"`
}
