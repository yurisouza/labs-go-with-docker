package services

import (
	"github.com/yurisouza/labs-go-with-docker/models"
	"github.com/yurisouza/labs-go-with-docker/repositories"
)

func GetAllUsers() ([]models.User, error) {

	users, err := repositories.GetAllUsers()

	if len(users) == 0 {
		return nil, err
	}

	return users, nil
}

func GetUser(id string) (*models.User, error) {
	return repositories.GetUser(id)
}

func InsertUser(user *models.User) error {
	return repositories.InsertUser(user)
}

func UpdateUser(user *models.User) error {
	return repositories.UpdateUser(user)
}

func RemoveUser(id string) error {
	return repositories.RemoveUser(id)
}

func CreateStructure() error {
	return repositories.CreateStructure()
}
