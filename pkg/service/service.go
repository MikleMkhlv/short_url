package service

import (
	"api"
	"api/pkg/repository"
)

type Short interface {
	Create(url api.Short) (string, error)
	GetUrl(shortUrl string) (string, error)
}

type Service struct {
	Short
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Short: NewShortService(repo.Short),
	}

}
