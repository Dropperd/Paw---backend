package service

import (
	"errors"
	"regexp"
	"strconv"
	"time"
	"websiteapi/entity"
	"websiteapi/entity/dto"
	"websiteapi/repository"
)

func GetUserImagesByClinicalId(clinicalId uint64, startDate time.Time, endDate time.Time, BodyPos string, UserId uint64) ([]entity.Image, error) {
	var images []entity.Image

	isClinical := repository.IsClinical(clinicalId)
	if !isClinical {
		return images, errors.New("user is not clinical")
	}

	images, err := repository.GetUserImagesByClinicalId(clinicalId, startDate, endDate, BodyPos, UserId)
	if err != nil {
		return images, err
	}

	if len(images) == 0 {
		return images, errors.New("no images found")
	}

	return images, nil
}

func GetMyImages(userId uint64, startDate time.Time, endDate time.Time, BodyPos string) ([]entity.Image, error) {
	var images []entity.Image

	images, err := repository.GetMyImages(userId, startDate, endDate, BodyPos)
	if err != nil {
		return images, err
	}

	if len(images) == 0 {
		return images, errors.New("no images found")
	}

	return images, nil
}

func GetImageById(imageId uint64) (entity.Image, error) {
	var image entity.Image

	image = repository.GetImageById(imageId)
	if image.ID == 0 {
		return image, errors.New("no image found")
	}

	return image, nil
}

func InsertImage(image entity.Image) (entity.Image, error) {

	boolExists := repository.BodyPositionExists(image.BodyPosition)
	if !boolExists {
		return image, errors.New("body position does not exist")
	}

	image = repository.InsertUpdateImageByID(image)

	return image, nil
}

func GetUsersForFilter(clinicalId uint64) []dto.UserIdDTO {
	return repository.GetUsersForFilter(clinicalId)
}

func UpdateImageById(image entity.Image, userId uint64) (entity.Image, error) {

	if image.UserID != userId {
		return image, errors.New("user does not own image")
	}

	image = repository.InsertUpdateImageByID(image)

	return image, nil

}

func DeleteImage(imageId uint64, userId uint64) error {
	var image entity.Image

	image = repository.GetImageById(imageId)
	if image.ID == 0 {
		return errors.New("no image found")
	}

	if image.UserID != userId {
		return errors.New("user does not own image")
	}

	repository.DeleteImage(imageId)

	return nil
}

func GetAllBodyPosition() ([]entity.BodyPos, error) {
	var bodyPositions []entity.BodyPos
	bodyPositions = repository.GetAllBodyPosition()

	if len(bodyPositions) == 0 {
		return bodyPositions, errors.New("no body positions found")
	}

	return bodyPositions, nil
}

func IsClinical(clinicalId uint64) bool {
	return repository.IsClinical(clinicalId)
}

func IsValidEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(regex, email)
	return match
}

func IsValidDate(date string) bool {
	layout := "2006-01-02"
	_, err := time.Parse(layout, date)
	return err == nil
}

func IsValidUserID(id string) bool {
	_, err := strconv.Atoi(id)
	return err == nil
}
