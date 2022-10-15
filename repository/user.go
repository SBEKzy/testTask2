package repository

import (
	"context"
	"github.com/SBEKzy/testTask2/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) CreateUser(u *model.User) error {
	userCollection := r.db.Collection("users")
	_, err := userCollection.InsertOne(context.TODO(), u)
	return err
}

func (r *repository) GetUser(email string) (*model.User, error) {
	var user model.User
	userCollection := r.db.Collection("users")
	err := userCollection.FindOne(context.TODO(), bson.D{{"email", email}}).Decode(&user)
	return &user, err
}

func (r *repository) UpdateUser(email string, update *model.User) error {
	userCollection := r.db.Collection("users")
	filter := bson.D{{"email", email}}
	_, err := userCollection.UpdateOne(context.TODO(), filter, bson.M{"$set": update})
	return err

}

func (r *repository) DeleteUser(email string) error {
	userCollection := r.db.Collection("users")
	filter := bson.D{{"email", email}}
	_, err := userCollection.DeleteOne(context.TODO(), filter)
	return err
}
