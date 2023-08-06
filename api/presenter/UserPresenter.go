package presenter

import (
	"learn_native/package/entites"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Username    string             `json:"username" bson:"username"`
	Email       string             `json:"email" bson:"email,omitempty"`
	PhoneNumber int                `json:"phone_number" bson:"phone_number,omitempty"`
	Address     string             `json:"address" bson:"address,omitempty"`
}

func UsersSuccessResponse(data *[]User) map[string]interface{} {
	response := map[string]interface{}{
		"statusCode": 200,
		"message":    "Success",
		"results":    data,
	}

	return response
}

func UserSuccessResponse(data *entites.Users, statusCode int) map[string]interface{} {
	user := User{
		ID:          data.ID,
		Name:        data.Name,
		Username:    data.Username,
		Email:       data.Email,
		PhoneNumber: data.PhoneNumber,
		Address:     data.Address,
	}

	response := map[string]interface{}{
		"statusCode": statusCode,
		"message":    "Success",
		"results":    user,
	}

	return response
}

func UserUpdateResponse(statusCode int) map[string]interface{} {
	response := map[string]interface{}{
		"statusCode": statusCode,
		"message":    "Success",
		"results":    nil,
		"error":      nil,
	}

	return response
}

func ErrorResponse(statusCode int, message string, errors error) map[string]interface{} {
	response := map[string]interface{}{
		"statusCode": statusCode,
		"message":    message,
		"results":    nil,
		"error":      errors.Error(),
	}

	return response
}
