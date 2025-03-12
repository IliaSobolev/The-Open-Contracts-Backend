package domain

import (
	"context"
	"github.com/pkg/errors"
)

type CodeBlock struct {
	ID          string `bson:"_id"`
	AuthorID    int64  `json:"author_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Rating      int64  `json:"rating"`
	Lang        string `json:"lang"`
	Body        string `json:"body"`
}

type CodeBlockDTO struct {
	AuthorID    int64  `json:"author_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Lang        string `json:"lang"`
	Body        string `json:"body"`
}

var ErrCodeBlockNotFound = errors.New("code block not found")

var Languages = []string{"func", "fift", "tact", "tolk"}

type CodeBlockUsecase interface {
	Create(ctx context.Context, codeBlock CodeBlockDTO) error
	List(ctx context.Context) ([]*CodeBlock, error)
	Get(ctx context.Context, id string) (*CodeBlock, error)
}

type CodeBlockRepository interface {
	Create(ctx context.Context, codeBlock *CodeBlock) error
	List(ctx context.Context) ([]*CodeBlock, error)
	Get(ctx context.Context, id string) (*CodeBlock, error)
}
