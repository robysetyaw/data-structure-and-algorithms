package services

import (
	"sda-learn/rest/models"
	"sda-learn/rest/repositories"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetAllUsers() []models.User {
	return s.Repo.GetAllUsers()
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	return s.Repo.GetUserByID(id)
}

func (s *UserService) InsertUser(user models.User) error {
	return s.Repo.InsertUser(user)
}

func (s *UserService) UpdateUser(user models.User) error {
	return s.Repo.UpdateUser(user)
}
