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

func (s *UserService) GetAll() ([]string, error) {
	result := []string{}
	for _, v := range s.data {
		result = append(result, v.Name)
	}
	return result, nil
}
