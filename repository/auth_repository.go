package repository

import (
	"websiteapi/config"
	"websiteapi/entity"
)

func Login(email string) (entity.User, error) {
	var user entity.User

	err := config.Db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
