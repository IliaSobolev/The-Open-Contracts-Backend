package usecase

import "TOC/pkg/domain"

type uc struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) domain.UserUsecase {
	return &uc{repo: repo}
}
