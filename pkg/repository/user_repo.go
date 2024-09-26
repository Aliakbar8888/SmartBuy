package repository

import (
	"errors"
	"products/db"
	"products/errs"
	"products/logger"
	"products/models"

	"gorm.io/gorm"
)

func CreateUser(user models.User) error {
	if err := db.GetDBConn().Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByID(userID uint) (models.User, error) {
	var user models.User
	if err := db.GetDBConn().Omit("password").First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, errs.ErrUserNotFound
		}
		return models.User{}, err
	}
	return user, nil
}

// func GetUserByUsername(username string) (user models.User, err error) {
// 	err = db.GetDBConn().Where("username=?", username).First(&user).Error
// 	if err != nil {
// 		logger.Error.Printf("[service.getuserbyusername]error in getting user by username  %s\n", err.Error())
// 		return user, translateErrors(err)
// 	}
// 	return user, nil
// }

func GetUserByUsernameAndPassword(username, password string) (user models.User, err error) {
	err = db.GetDBConn().Where("username=? AND password=?", username, password).First(&user).Error
	if err != nil {
		logger.Error.Printf("[service.getuserbyusername]error in getting user by username  %s\n", err.Error())

		return user, translateErrors(err)
	}
	return user, nil
}

func UpdateUser(user models.User) error {
	if err := db.GetDBConn().Save(user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(userID uint) error {
	if err := db.GetDBConn().Delete(&models.User{}, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.ErrUserNotFound
		}
		return err
	}
	return nil
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := db.GetDBConn().Omit("password").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
