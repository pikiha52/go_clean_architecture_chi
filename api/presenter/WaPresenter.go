package presenter

func SuccessResponse(statusCode int, message string) map[string]interface{} {
	response := map[string]interface{}{
		"statusCode": statusCode,
		"message":    message,
		"resutls":    nil,
		"error":      nil,
	}

	return response
}
