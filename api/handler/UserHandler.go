package handler

import (
	"encoding/json"
	"learn_native/api/presenter"
	"learn_native/package/users"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

		response := presenter.UserSuccessResponse(requests, 201)
		encode, _ := json.Marshal(response)

		w.WriteHeader(201)
		w.Write(encode)
	}
}

func ShowHandler(service users.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id := chi.URLParam(r, "id")
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			responseError := presenter.ErrorResponse(500, "Gagal transform string ke objectID!", err)
			encodeError, _ := json.Marshal(responseError)
			w.WriteHeader(500)
			w.Write(encodeError)
			return
		}

		requests, err := service.ShowService(objID)
		if err != nil {
			responseErrorService := presenter.ErrorResponse(404, "Not found!", err)
			encodeErrorService, _ := json.Marshal(responseErrorService)
			w.WriteHeader(404)
			w.Write(encodeErrorService)
			return
		}

		respone := presenter.UserSuccessResponse(requests, 200)

		encode, _ := json.Marshal(respone)

		w.WriteHeader(200)
		w.Write(encode)
	}
}

func UpdateHandler(service users.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id := chi.URLParam(r, "id")
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			responseError := presenter.ErrorResponse(500, "Gagal transform string ke objectID!", err)
			encodeError, _ := json.Marshal(responseError)
			w.WriteHeader(500)
			w.Write(encodeError)
			return
		}

		err = service.UpdateService(objID, r)
		if err != nil {
			w.Write([]byte("Error service!"))
		}

		response := presenter.UserUpdateResponse(200)
		encode, _ := json.Marshal(response)

		w.WriteHeader(200)
		w.Write(encode)
	}
}
