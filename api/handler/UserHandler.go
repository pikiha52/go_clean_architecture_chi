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
			w.Write([]byte("Failed get data"))
		}

		response := presenter.UsersSuccessResponse(requests)
		decode, err := json.Marshal(response)
		if err != nil {
			w.Write([]byte("Failed get data"))
		}
		w.Write(decode)

	}
}
