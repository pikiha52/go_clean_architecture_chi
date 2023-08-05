package users

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	TestRepository() error
}

type repository struct {
	Collection *mongo.Collection
}

func NewRepo(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}

func (r *repository) TestRepository() error {
	return nil
}
