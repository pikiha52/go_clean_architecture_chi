package users

import (
	"context"
	"errors"
	"learn_native/api/contract"
	"learn_native/api/presenter"
	"learn_native/package/entites"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	IndexRepository() (*[]presenter.User, error)
	StoreRepository(contractCreate contract.UserCreate) (*entites.Users, error)
	ShowRepository(id primitive.ObjectID) (*entites.Users, error)
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

func (r *repository) StoreRepository(contractCreate contract.UserCreate) (*entites.Users, error) {
	var user entites.Users
	user.ID = primitive.NewObjectID()
	user.Name = contractCreate.Name
	user.Username = contractCreate.Username
	user.Email = contractCreate.Email
	user.PhoneNumber = contractCreate.PhoneNumber
	user.Address = contractCreate.Address

	_, err := r.Collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) ShowRepository(id primitive.ObjectID) (*entites.Users, error) {
	var user entites.Users
	err := r.Collection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: id}}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("User tidak ditemukan!")
		}
	}

	return &user, nil
}
