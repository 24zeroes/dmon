package service

import "github.com/24zeroes/dmon/pkg/repository"

type Authorization interface {
}

type DmonList interface {
}

type DmonItem interface {
}

type Service struct {
	Authorization
	DmonList
	DmonItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
