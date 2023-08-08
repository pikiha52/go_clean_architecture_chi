package main

import (
	"context"
	"fmt"
	"learn_native/api/routes"
	"learn_native/package/users"
	"learn_native/package/wa"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	chiRouter := chi.NewRouter()
	chiRouter.Use(middleware.Logger)

	db, _, err := DatabaseConnection()

	if err != nil {
		fmt.Println("Gagal menghubungkan ke MongoDB!")
	}

	collection := db.Collection("users")
	collectionOtp := db.Collection("otp")

	userRepo := users.NewRepo(collection)
	waRepo := wa.NewRepo("https://waba-sandbox.360dialog.io/", collectionOtp)
	waService := wa.NewService(waRepo)
	userService := users.NewService(userRepo, waService)

	chiRouter.Route("/api", func(r chi.Router) {
		routes.SetupRouteUser(r, userService)
	})

	http.ListenAndServe(":3000", chiRouter)
}

func DatabaseConnection() (*mongo.Database, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb://localhost:27017/go_db").SetServerSelectionTimeout(5*time.
		Second))

	if err != nil {
		cancel()
		return nil, nil, err
	}

	db := client.Database("go_db")
	return db, cancel, nil
}
