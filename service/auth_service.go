package service

import (
	"websiteapi/entity/dto"
	"websiteapi/repository"

	"golang.org/x/crypto/bcrypt"
)

func Login(loginDTO dto.LoginDTO) (string, error) {

	token := ""

	user, err := repository.Login(loginDTO.Email)
	if err != nil {
		return token, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDTO.Password))
	if err != nil {
		return token, err
	}

	token, err = CreateToken(user.ID)
	if err != nil {
		return token, err
	}

	return token, nil
}
