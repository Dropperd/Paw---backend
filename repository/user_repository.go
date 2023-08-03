package repository

import (
	"errors"
	"websiteapi/config"
	"websiteapi/entity"
)

func GetAlluser() []entity.User {
	var user []entity.User
	config.Db.Table("users").Find(&user)
	return user
}

func GetSharedClinicals(userId uint64) ([]entity.User, error) {
	var clinicals []entity.User

	err := config.Db.Raw("SELECT u.name,u.id,u.email FROM users u, image_clinicals ic WHERE ic.user_id = ? AND ic.clinical_id = u.id AND u.user_type = 1", userId).Scan(&clinicals).Error
	if err != nil {
		return nil, err
	}
	return clinicals, nil
}

func GetUser(userID uint64) (entity.User, error) {
	var user entity.User
	config.Db.First(&user, userID)
	if user.ID != 0 {
		return user, nil
	}
	return user, errors.New("user not found")
}

func InsertUpdateUserByID(user entity.User) entity.User {
	var existingUser entity.User
	config.Db.Where("id = ?", user.ID).First(&existingUser)
	if existingUser.ID != 0 {
		if existingUser.ID == user.ID {
			config.Db.Model(&existingUser).UpdateColumns(user)
		}
	} else {
		config.Db.Save(&user)
	}
	return existingUser
}

func DeleteUserByID(id uint64) error {
	var user entity.User

	config.Db.Table("users").First(&user, id)
	if user.ID != 0 {
		config.Db.Table("users").Delete(&user)
		return nil
	}
	return errors.New("user not found")

}

func GetUserType() any {
	var userType []struct {
		Id   int    `json:"id"`
		Type string `json:"type"`
	}
	config.Db.Raw("SELECT * FROM user_type").Scan(&userType)
	return userType
}

func ExistsEmail(email string) bool {
	var user entity.User
	config.Db.Where("email = ?", email).First(&user)
	return user.ID != 0
}

func GetAllClinical() []entity.User {
	var clinicals []entity.User
	config.Db.Table("users").Where("user_type = ?", 1).Find(&clinicals)
	return clinicals
}

func InsertImageClinical(imageClinical entity.ImageClinical) {
	var existingImageClinical entity.ImageClinical
	config.Db.Table("image_clinicals").Where("user_id = ? AND clinical_id = ?", imageClinical.UserId, imageClinical.ClinicalId).First(&existingImageClinical)
	if existingImageClinical.ID == 0 {
		config.Db.Table("image_clinicals").Save(&imageClinical)
	}
}
