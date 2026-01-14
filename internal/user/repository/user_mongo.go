package repository

import (
	"TOC/pkg/domain"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	c *mongo.Collection
}

func NewUserRepository(db *mongo.Database) domain.UserRepository {
	return &repo{db.Collection("users")}
}

func (r repo) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r repo) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}
