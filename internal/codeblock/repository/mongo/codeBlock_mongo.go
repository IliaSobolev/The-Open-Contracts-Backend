package mongo

import (
	"TOC/pkg/domain"
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	c *mongo.Collection
}

func NewCodeblockRepo(db *mongo.Database) domain.CodeBlockRepository {
	return &repo{db.Collection("codeblocks")}
}

func (r repo) Create(ctx context.Context, codeBlock *domain.CodeBlock) error {
	_, err := r.c.InsertOne(ctx, codeBlock)
	if err != nil {
		return errors.Wrap(err, "mongo")
	}
	return nil
}

func (r repo) Get(ctx context.Context, id string) (*domain.CodeBlock, error) {
	res := r.c.FindOne(ctx, bson.M{"_id": id})
	if errors.Is(res.Err(), mongo.ErrNoDocuments) {
		return nil, domain.ErrCodeBlockNotFound
	} else if res.Err() != nil {
		return nil, errors.Wrap(res.Err(), "mongo")
	}
	var obj domain.CodeBlock
	err := res.Decode(&obj)
	if err != nil {
		return nil, errors.Wrap(err, "mongo")
	}
	return &obj, nil
}

func (r repo) List(ctx context.Context) ([]*domain.CodeBlock, error) {
	res, err := r.c.Find(ctx, bson.D{})
	if err != nil {
		return nil, domain.ErrCodeBlockNotFound
	}
	var codeBlocks []*domain.CodeBlock
	if err := res.All(ctx, &codeBlocks); err != nil {
		return nil, domain.ErrCodeBlockNotFound
	}
	return codeBlocks, nil
}
