package service

import (
	"github.com/SBEKzy/testTask2/model"
	"github.com/gin-gonic/gin"
)

type Repository interface {
	CreateUser(c *gin.Context, u *model.User) error
	GetUser(c *gin.Context, email string) (*model.User, error)
	UpdateUser(c *gin.Context, email string, update *model.User) error
	DeleteUser(c *gin.Context, email string) error
}

type service struct {
	repo Repository
}

func New(r Repository) *service {
	return &service{repo: r}
}
