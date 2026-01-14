package domain

import "context"

type User struct {
	ID        int64  `json:"id" bson:"id"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Username  string `json:"username" bson:"username"`
	PhotoURL  string `json:"photo_url" bson:"photo_url"`
	AuthDate  string `json:"auth_date" bson:"auth_date"`
	Hash      string `json:"hash" bson:"hash"`
}

type UserUsecase interface {
	Create(ctx context.Context, user *User) (*User, error)
	GetByID(ctx context.Context, id int64) (*User, error)
	Authenticate(ctx context.Context, user *User) (bool, error)
}

type UserRepository interface {
	Create(ctx context.Context, user *User) (*User, error)
	GetByID(ctx context.Context, id int64) (*User, error)
}
