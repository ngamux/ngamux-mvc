package model

import "go.mongodb.org/mongo-driver/mongo"

type Model struct {
	db *mongo.Database
}

func New(db *mongo.Database) *Model {
	return &Model{
		db: db,
	}
}
