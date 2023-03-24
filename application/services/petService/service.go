package service

import (
	"github.com/scarlettmiss/bestPal/application/domain/pet"
)

type Service struct {
	repo pet.Repository
}

func New(repo pet.Repository) (Service, error) {
	return Service{repo: repo}, nil
}
