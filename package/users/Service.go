package users

import (
	"encoding/json"
	"learn_native/api/contract"
	"learn_native/api/presenter"
	"learn_native/package/entites"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	IndexService() (*[]presenter.User, error)
	StoreService(httpRequest *http.Request) (*entites.Users, error)
	ShowService(id primitive.ObjectID) (*entites.Users, error)
	UpdateService(id primitive.ObjectID, httpRequest *http.Request) error
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) IndexService() (*[]presenter.User, error) {
	return s.repository.IndexRepository()
}

func (s *service) StoreService(httpRequest *http.Request) (*entites.Users, error) {
	var contractCreate contract.UserCreate

	err := json.NewDecoder(httpRequest.Body).Decode(&contractCreate)

	if err != nil {
		return nil, err
	}

	data, err := s.repository.StoreRepository(contractCreate)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *service) ShowService(id primitive.ObjectID) (*entites.Users, error) {
	data, err := s.repository.ShowRepository(id)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *service) UpdateService(id primitive.ObjectID, httpRequest *http.Request) error {
	var contractUpdate contract.UserUpdate
	err := json.NewDecoder(httpRequest.Body).Decode(&contractUpdate)

	if err != nil {
		return err
	}

	err = s.repository.UpdateRepository(id, contractUpdate)
	if err != nil {
		return err
	}

	return nil
}
