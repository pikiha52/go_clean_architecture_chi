package users

import (
	"encoding/json"
	"errors"
	"fmt"
	"learn_native/api/contract"
	"learn_native/api/presenter"
	"learn_native/package/entites"
	"learn_native/package/wa"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	IndexService() (*[]presenter.User, error)
	StoreService(httpRequest *http.Request) (*entites.Users, error)
	ShowService(id primitive.ObjectID) (*entites.Users, error)
	UpdateService(id primitive.ObjectID, httpRequest *http.Request) error
	UserOtpService(id primitive.ObjectID, httpRequest *http.Request) error
}

type service struct {
	repository Repository
	waService  wa.Service
}

func NewService(repo Repository, waService wa.Service) Service {
	return &service{
		repository: repo,
		waService:  waService,
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

	s.waService.SendWA(*data, "")

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

func (s *service) UserOtpService(id primitive.ObjectID, httpRequest *http.Request) error {
	var otpContract contract.OtpCheck
	err := json.NewDecoder(httpRequest.Body).Decode(&otpContract)

	if err != nil {
		return err
	}

	user, err := s.ShowService(id)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("User not found!")
	}

	_, err = s.waService.CheckOtpService(otpContract)
	if err != nil {
		return err
	}

	sendWa := s.waService.SendWA(*user, "Selamat, akun anda sudah aktif!")
	fmt.Println(sendWa)
	if sendWa != nil {
		return sendWa
	}

	return nil
}
