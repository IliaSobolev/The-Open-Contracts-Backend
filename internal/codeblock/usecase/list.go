package usecase

import (
	"TOC/pkg/domain"
	"context"
	"github.com/pkg/errors"
)

func (uc *uc) List(ctx context.Context) ([]*domain.CodeBlock, error) {
	res, err := uc.repo.List(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "codeBlock repo.List")
	}
	return res, nil
}
