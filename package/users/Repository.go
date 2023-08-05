package users

import (
	"context"
	"learn_native/api/presenter"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	IndexRepository() (*[]presenter.User, error)
}

type repository struct {
	Collection *mongo.Collection
}

func NewRepo(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}

func (r *repository) IndexRepository() (*[]presenter.User, error) {
	var users []presenter.User
	cursor, err := r.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var user presenter.User
		_ = cursor.Decode(&user)
		users = append(users, user)
	}

	return &users, nil
}
