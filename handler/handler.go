package handler

import "github.com/SBEKzy/testTask2/model"

type Service interface {
	CreateUser(u *model.User) error
	GetUser(email string) (*model.User, error)
	UpdateUser(email string, update *model.User) error
	DeleteUser(email string) error
}

type handler struct {
	service Service
}

func New(service Service) *handler {
	return &handler{service: service}
}
