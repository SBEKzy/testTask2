package service

import (
	"github.com/SBEKzy/testTask2/model"
	"github.com/gin-gonic/gin"
)

func (s *service) CreateUser(c *gin.Context, u *model.User) error {
	err := s.repo.CreateUser(c, u)
	return err
}

func (s *service) GetUser(c *gin.Context, email string) (*model.User, error) {
	user, err := s.repo.GetUser(c, email)
	return user, err
}

func (s *service) UpdateUser(c *gin.Context, email string, update *model.User) error {
	err := s.repo.UpdateUser(c, email, update)
	return err
}

func (s *service) DeleteUser(c *gin.Context, email string) error {
	err := s.repo.DeleteUser(c, email)
	return err
}
