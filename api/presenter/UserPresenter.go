package presenter

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID          primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Username    string             `json:"username" bson:"username"`
	Email       string             `json:"email" bson:"email"`
	PhoneNumber int                `json:"phone_number" bson:"phone_number"`
	Address     string             `json:"address" bson:"address"`
}

func UsersSuccessResponse(data *[]User) map[string]interface{} {
	response := map[string]interface{}{
		"statusCode": 200,
		"message":    "Success",
		"results":    data,
	}

	return response
}

func ErrorResponse(statusCode int, message string, errors error) map[string]interface{} {
	response := map[string]interface{}{
		"statusCode": statusCode,
		"message":    message,
		"results":    nil,
		"error":      errors,
	}

	return response
}
