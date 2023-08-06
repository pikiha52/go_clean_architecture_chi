package routes

import (
	"learn_native/api/handler"
	"learn_native/package/users"

	"github.com/go-chi/chi/v5"
)

func SetupRouteUser(chiRouter chi.Router, service users.Service) {
	chiRouter.Get("/users", handler.IndexHandler(service))
	chiRouter.Post("/user", handler.StoreHandler(service))
	chiRouter.Get("/user/{id}", handler.ShowHandler(service))
	chiRouter.Put("/user/{id}", handler.UpdateHandler(service))
}
