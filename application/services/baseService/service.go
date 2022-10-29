package service

import "github.com/scarlettmiss/bestPal/application/repositories/baseRepo"

type Service struct {
	repo *baseRepo.Repository
}

func New(repo *baseRepo.Repository) (*Service, error) {
	return &Service{repo: repo}, nil
}

/**
implements the domains methods for service so application can call them
*/
