package contract

type UserCreate struct {
	Name        string `json:"name" bson:"name"`
	Username    string `json:"username" bson:"username"`
	Email       string `json:"email" bson:"email"`
	PhoneNumber int    `json:"phone_number" bson:"phone_number"`
	Address     string `json:"address" bson:"address"`
}

type UserUpdate struct {
	Name        string `json:"name" bson:"name"`
	Username    string `json:"username" bson:"username"`
	Email       string `json:"email" bson:"email"`
	PhoneNumber int    `json:"phone_number" bson:"phone_number"`
	Address     string `json:"address" bson:"address"`
}

type OtpCheck struct {
	OtpCode int `json:"otp"`
}
