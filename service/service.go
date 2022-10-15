package service

import "github.com/SBEKzy/testTask2/model"

type Repository interface {
	CreateUser(u *model.User) error
	GetUser(email string) (*model.User, error)
	UpdateUser(email string, update *model.User) error
	DeleteUser(email string) error
}

type service struct {
	repo Repository
}

func New(r Repository) *service {
	return &service{repo: r}
}
