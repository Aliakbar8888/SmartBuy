package service

import (
	"fmt"
	"products/models"
	"products/pkg/repository"
	"products/utils"
)

func GetAllUsers() (users []models.User, err error) {
	users, err = repository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserByID(id uint) (user models.User, err error) {
	user, err = repository.GetUserByID(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func CreateUser(user models.User) error {
	_, err := repository.GetUserByUsernameAndPassword(user.Username, user.Password)
	if err != nil {
		fmt.Println(err)
	}
	user.Password = utils.GenerateHash(user.Password)

	err = repository.CreateUser(user)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
