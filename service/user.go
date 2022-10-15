package service

import "github.com/SBEKzy/testTask2/model"

func (s *service) CreateUser(u *model.User) error {
	err := s.repo.CreateUser(u)
	return err
}

func (s *service) GetUser(email string) (*model.User, error) {
	user, err := s.repo.GetUser(email)
	return user, err
}

func (s *service) UpdateUser(email string, update *model.User) error {
	err := s.repo.UpdateUser(email, update)
	return err
}

func (s *service) DeleteUser(email string) error {
	err := s.repo.DeleteUser(email)
	return err
}
