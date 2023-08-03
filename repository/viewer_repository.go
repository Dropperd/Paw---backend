package repository

import (
	"errors"
	"time"
	"websiteapi/config"
	"websiteapi/entity"
	"websiteapi/entity/dto"
)

func IsClinical(userId uint64) bool {
	var user entity.User
	config.Db.Table("users").Where("id = ?", userId).First(&user)
	return user.UserType == 1
}

func GetUserImagesByClinicalId(clinicalId uint64, startDate time.Time, endDate time.Time, BodyPos string, UserId uint64) ([]entity.Image, error) {
	if !startDate.IsZero() && !endDate.IsZero() && startDate.After(endDate) {
		return nil, errors.New("invalid date range. Start date cannot be higher than end date")
	}

	var images []entity.Image
	query := config.Db.Table("image").Model(&entity.Image{}).
		Joins("JOIN image_clinicals ic ON image.user_id = ic.user_id").
		Where("ic.clinical_id = ?", clinicalId)

	if UserId != 0 {
		query = query.Where("ic.user_id = ?", UserId).Where("ic.clinical_id = ?", clinicalId)
	}

	if !startDate.IsZero() && !endDate.IsZero() {
		query = query.Where("updated_at BETWEEN ? AND ?", startDate, endDate)
	}

	if BodyPos != "" {
		if !(BodyPos == "none") {
			query = query.Where("body_position = ?", BodyPos)
		}
	}

	if err := query.Find(&images).Error; err != nil {
		return nil, err
	}

	return images, nil
}

func GetMyImages(userId uint64, startDate time.Time, endDate time.Time, BodyPos string) ([]entity.Image, error) {
	if !startDate.IsZero() && !endDate.IsZero() && startDate.After(endDate) {
		return nil, errors.New("invalid date range. Start date cannot be higher than end date")
	}

	var images []entity.Image
	query := config.Db.Table("image").Model(&entity.Image{}).
		Where("user_id = ?", userId)

	if !startDate.IsZero() && !endDate.IsZero() {
		query = query.Where("updated_at BETWEEN ? AND ?", startDate, endDate)
	}

	if BodyPos != "" {
		query = query.Where("body_position = ?", BodyPos)
	}

	if err := query.Find(&images).Error; err != nil {
		return nil, err
	}

	return images, nil
}

func GetImageById(imageId uint64) entity.Image {
	var image entity.Image
	config.Db.Table("image").Where("id = ?", imageId).First(&image)
	return image
}

func GetUsersForFilter(clinicalId uint64) []dto.UserIdDTO {
	var users []dto.UserIdDTO
	config.Db.Table("image_clinicals").Model(&dto.UserIdDTO{}).Where("clinical_id = ?", clinicalId).Find(&users)
	return users
}

func InsertUpdateImageByID(image entity.Image) entity.Image {
	var existingImage entity.Image
	config.Db.Table("image").Where("id = ?", image.ID).First(&existingImage)
	if existingImage.ID != 0 {
		if existingImage.ID == image.ID {
			image.Updated_At = time.Now().Format("2006-01-02")
			config.Db.Table("image").Model(&existingImage).UpdateColumns(image)
		}
	} else {
		image.Added_At = time.Now().Format("2006-01-02")
		image.Updated_At = time.Now().Format("2006-01-02")
		config.Db.Table("image").Save(&image)
	}
	return image
}

func DeleteImage(image uint64) {
	config.Db.Raw("DELETE FROM image WHERE id = ?", image).Scan(&image)
}

func GetAllBodyPosition() []entity.BodyPos {
	var bodyPositions []entity.BodyPos
	config.Db.Table("body_pos").Find(&bodyPositions)
	return bodyPositions
}

func BodyPositionExists(bodyPos string) bool {
	var bodyPosition entity.BodyPos
	config.Db.Table("body_pos").Where("membro = ?", bodyPos).First(&bodyPosition)
	return bodyPosition.Id != 0
}
