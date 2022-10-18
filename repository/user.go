package repository

import (
	"github.com/SBEKzy/testTask2/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) CreateUser(c *gin.Context, u *model.User) error {
	userCollection := r.db.Collection("users")
	_, err := userCollection.InsertOne(c, u)
	return err
}

func (r *repository) GetUser(c *gin.Context, email string) (*model.User, error) {
	var user model.User
	userCollection := r.db.Collection("users")
	err := userCollection.FindOne(c, bson.D{{"email", email}}).Decode(&user)
	return &user, err
}

func (r *repository) UpdateUser(c *gin.Context, email string, update *model.User) error {
	userCollection := r.db.Collection("users")
	filter := bson.D{{"email", email}}
	_, err := userCollection.UpdateOne(c, filter, bson.M{"$set": update})
	return err

}

func (r *repository) DeleteUser(c *gin.Context, email string) error {
	userCollection := r.db.Collection("users")
	filter := bson.D{{"email", email}}
	_, err := userCollection.DeleteOne(c, filter)
	return err
}
