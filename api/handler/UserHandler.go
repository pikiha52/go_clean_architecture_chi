package handler

import (
	"encoding/json"
	"learn_native/api/presenter"
	"learn_native/package/users"
	"net/http"
)

func IndexHandler(service users.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		requests, err := service.IndexService()
		if err != nil {
			errorResponse := presenter.ErrorResponse(500, "Error!", err)
			errorDecode, _ := json.Marshal(errorResponse)
			w.WriteHeader(500)
			w.Write(errorDecode)
			return
		}

		response := presenter.UsersSuccessResponse(requests)
		decode, err := json.Marshal(response)
		if err != nil {
			errorResponse := presenter.ErrorResponse(500, "Error!", err)
			errorDecode, _ := json.Marshal(errorResponse)
			w.WriteHeader(500)
			w.Write(errorDecode)
			return
		}

		w.WriteHeader(200)
		w.Write(decode)
	}
}

func StoreHandler(service users.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		requests, err := service.StoreService(r)
		if err != nil {
			errResponse := presenter.ErrorResponse(500, "Error", err)
			errEncode, _ := json.Marshal(errResponse)
			w.WriteHeader(500)
			w.Write(errEncode)
			return
		}

		response := presenter.UserSuccessResponse(requests)
		encode, _ := json.Marshal(response)

		w.WriteHeader(201)
		w.Write(encode)
	}
}
