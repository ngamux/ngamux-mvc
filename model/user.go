package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
}

func (m *Model) GetAllUsers(ctx context.Context) ([]User, error) {
	q, err := m.db.Collection("users").Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	users := make([]User, 0)
	if err = q.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (m *Model) CreateUser(ctx context.Context, user User) (*User, error) {
	_, err := m.db.Collection("users").InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
