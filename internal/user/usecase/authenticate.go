package usecase

import (
	"TOC/pkg/domain"
	"context"
)

func (uc *uc) Authenticate(ctx context.Context, user *domain.User) (bool, error) {
	return true, nil
}
