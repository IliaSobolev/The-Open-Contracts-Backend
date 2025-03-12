package usecase

import (
	"TOC/pkg/domain"
	"context"
	"github.com/pkg/errors"
)

func (uc *uc) Get(ctx context.Context, id string) (*domain.CodeBlock, error) {
	res, err := uc.repo.Get(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "codeBlock repo.Get")
	}
	return res, nil
}
