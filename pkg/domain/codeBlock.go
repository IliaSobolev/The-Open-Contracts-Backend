package domain

import (
	"context"
	"github.com/pkg/errors"
)

type CodeBlock struct {
	ID          string `json:"_id" bson:"_id"`
	AuthorName  string `json:"author_name" bson:"author_name"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Rating      int64  `json:"rating" bson:"rating"`
	Lang        string `json:"lang" bson:"lang"`
	Body        string `json:"body" bson:"body"`
}

type CodeBlockDTO struct {
	AuthorName  string `json:"author_name" bson:"author_name"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Lang        string `json:"lang" bson:"lang"`
	Body        string `json:"body" bson:"body"`
}

var ErrCodeBlockNotFound = errors.New("code block not found")
var ErrInvalidLanguageCode = errors.New("invalid language")

var Languages = []string{"func", "fift", "tact", "tolk"}

type CodeBlockUsecase interface {
	Create(ctx context.Context, codeBlock CodeBlockDTO) error
	List(ctx context.Context, authorName string, lang string, sortOrder string) ([]*CodeBlock, error)
	Get(ctx context.Context, id string) (*CodeBlock, error)
}

type CodeBlockRepository interface {
	Create(ctx context.Context, codeBlock *CodeBlock) error
	List(ctx context.Context, params SearchParams) ([]*CodeBlock, error)
	Get(ctx context.Context, id string) (*CodeBlock, error)
}
