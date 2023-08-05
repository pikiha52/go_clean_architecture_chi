package users

import "learn_native/api/presenter"

type Service interface {
	IndexService() (*[]presenter.User, error)
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
	users := []presenter.User{
		{
			ID:          1,
			Name:        "John doe",
			Username:    "johndoe",
			Email:       "johndoe@dev.com",
			PhoneNumber: 628829100012,
			Address:     "Jakarta, Indonesia",
		},
	}

	return &users, nil
}
