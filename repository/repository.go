package repository

import "go.mongodb.org/mongo-driver/mongo"

type repository struct {
	db *mongo.Database
}

func New(db *mongo.Database) *repository {
	//collection := db.Collection("users")
	return &repository{db: db}
}
