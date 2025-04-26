package usecase

import (
	"TOC/pkg/domain"
	"TOC/pkg/utils"
	"context"
	"github.com/pkg/errors"
)

func (uc *uc) List(ctx context.Context, authorName string, lang string, sortOrder string) ([]*domain.CodeBlock, error) {
	params := domain.SearchParams{SortOrder: sortOrder}
	if lang != "" {
		isValidLang := utils.LangValidation(lang)
		if !isValidLang {
			return nil, domain.ErrInvalidLanguageCode
		}
		params.Lang = &lang
	}
	if authorName != "" {
		params.AuthorName = &authorName
	}

	res, err := uc.repo.List(ctx, params)
	if err != nil {
		return nil, errors.Wrap(err, "codeBlock repo.List")
	}
	return res, nil
}
