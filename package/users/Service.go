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
	return s.repository.IndexRepository()
}
