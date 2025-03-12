package usecase

import (
	"TOC/pkg/domain"
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"regexp"
)

var titleRegex = regexp.MustCompile("^[A-Za-z ]+$")

func (uc *uc) Create(ctx context.Context, codeBlockDTO domain.CodeBlockDTO) error {
	if codeBlockDTO.Title == "" || codeBlockDTO.Description == "" || codeBlockDTO.Body == "" || !titleRegex.MatchString(codeBlockDTO.Title) {
		return errors.New("invalid data")
	}
	found := false
	for _, lang := range domain.Languages {
		if lang == codeBlockDTO.Lang {
			found = true
			break
		}
	}
	if !found {
		return errors.New("invalid data")
	}
	codeBlock := domain.CodeBlock{
		ID:          uuid.NewString(),
		AuthorID:    codeBlockDTO.AuthorID,
		Title:       codeBlockDTO.Title,
		Description: codeBlockDTO.Description,
		Rating:      0,
		Lang:        codeBlockDTO.Lang,
		Body:        codeBlockDTO.Body,
	}
	err := uc.repo.Create(ctx, &codeBlock)
	if err != nil {
		return errors.Wrap(err, "codeBlock repo.Create")
	}
	return nil
}
