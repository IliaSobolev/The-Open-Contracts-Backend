package usecase

import "TOC/pkg/domain"

type uc struct {
	repo domain.CodeBlockRepository
}

func NewCodeblockUsecase(repo domain.CodeBlockRepository) domain.CodeBlockUsecase {
	return &uc{repo: repo}
}
