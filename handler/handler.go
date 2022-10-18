package handler

import (
	"github.com/SBEKzy/testTask2/model"
	"github.com/gin-gonic/gin"
)

type Service interface {
	CreateUser(c *gin.Context, u *model.User) error
	GetUser(c *gin.Context, email string) (*model.User, error)
	UpdateUser(c *gin.Context, email string, update *model.User) error
	DeleteUser(c *gin.Context, email string) error
}

type handler struct {
	service Service
}

func New(service Service) *handler {
	return &handler{service: service}
}
