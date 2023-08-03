package service

import (
	"errors"
	"log"
	"websiteapi/entity"
	"websiteapi/entity/dto"
	"websiteapi/repository"

	"golang.org/x/crypto/bcrypt"

	"github.com/mashingan/smapping"
)

func GetAllClinical() []entity.User {
	var clinicals []entity.User

	clinicals = repository.GetAllClinical()

	return clinicals
}

func GetSharedClinicals(userId uint64) ([]entity.User, error) {
	var clinicals []entity.User

	clinicals, err := repository.GetSharedClinicals(userId)

	if err != nil {
		return nil, err
	}

	return clinicals, nil
}

func InsertImageClinical(clinical entity.ImageClinical) {
	repository.InsertImageClinical(clinical)
}

func GetAlluser() []dto.UserDTO {
	usersResponse := []dto.UserDTO{}
	var users = repository.GetAlluser()

	for _, user := range users {
		userResponse := dto.UserDTO{}
		err := smapping.FillStruct(&userResponse, smapping.MapFields(user))
		if err != nil {
			log.Println("failed to map to response ", err)
		}
		usersResponse = append(usersResponse, userResponse)
	}
	return usersResponse
}

func InsertUser(user entity.User) (dto.UserCreatedorUpdatedDTO, error) {
	userCreated := dto.UserCreatedorUpdatedDTO{}

	if repository.ExistsEmail(user.Email) {
		return userCreated, errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("failed to hash password", err)
		return userCreated, err
	}

	user.Password = string(hashedPassword)

	userReturn := repository.InsertUpdateUserByID(user)
	err = smapping.FillStruct(&userCreated, smapping.MapFields(&userReturn))
	if err != nil {
		log.Fatal("failed map", err)
		return userCreated, err
	}

	return userCreated, nil
}

func GetUserProfile(userID uint64) (dto.UserDTO, error) {
	userResponse := dto.UserDTO{}

	user, err1 := repository.GetUser(userID)
	if err1 != nil {
		log.Println("failed to get user ", err1)
		return userResponse, err1
	}
	err := smapping.FillStruct(&userResponse, smapping.MapFields(&user))
	if err != nil {
		log.Println("failed to map to response ", err)
		return userResponse, err
	}

	return userResponse, nil
}

func UpdateUserByID(user entity.User, userID uint64) (dto.UserCreatedorUpdatedDTO, error) {
	userUpdated := dto.UserCreatedorUpdatedDTO{}
	if user.ID != userID {
		return userUpdated, errors.New("you can't update this user")
	}
	userReturn := repository.InsertUpdateUserByID(user)
	err := smapping.FillStruct(&userUpdated, smapping.MapFields(&userReturn))
	if err != nil {
		log.Fatal("failed map update user by id", err)
	}
	return userUpdated, nil
}

func DeleteUserByID(id uint64, userID uint64) error {
	if id != userID {
		return errors.New("you can't delete this user")
	}

	return repository.DeleteUserByID(id)
}

func GetUserType() any {
	return repository.GetUserType()
}
