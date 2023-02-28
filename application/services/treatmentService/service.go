package service

import (
	"github.com/scarlettmiss/bestPal/application/domain/treatment"
)

type Service struct {
	repo treatment.Repository
}

func New(repo treatment.Repository) (Service, error) {
	return Service{repo: repo}, nil
}
