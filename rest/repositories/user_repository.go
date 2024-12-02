package repositories

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"sda-learn/rest/models"
)

type UserRepository struct {
	Users []models.User
}

func NewUserRepository(filePath string) (*UserRepository, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var users []models.User
	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}

	return &UserRepository{Users: users}, nil
}

func (r *UserRepository) GetAllUsers() []models.User {
	return r.Users
}

func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	for _, user := range r.Users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *UserRepository) InsertUser(newUser models.User) error {
	// Check if ID already exists
	for _, user := range r.Users {
		if user.ID == newUser.ID {
			return errors.New("user with this ID already exists")
		}
	}
	r.Users = append(r.Users, newUser)
	return nil
}

func (r *UserRepository) UpdateUser(updatedUser models.User) error {
	for i, user := range r.Users {
		if user.ID == updatedUser.ID {
			r.Users[i] = updatedUser
			return nil
		}
	}
	return errors.New("user not found")
}