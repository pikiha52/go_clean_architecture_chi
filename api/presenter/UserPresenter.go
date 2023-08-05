package presenter

type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber int    `json:"phone_number"`
	Address     string `json:"address"`
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
