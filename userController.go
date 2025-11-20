package main

import (
	"project-test/models"
)

type UserService struct {
	data map[string]*models.User
}

func NewUserService() *UserService {
	return &UserService{
		data: make(map[string]*models.User),
	}
}

func (s *UserService) GetAll() ([]*models.User, error) {
	result := []*models.User{}
	for _, v := range s.data {
		result = append(result, v)
	}
	return result, nil
}
