package usecase

import (
	entity "immersiveProject/features/login/entity"
	"immersiveProject/middlewares"

	"golang.org/x/crypto/bcrypt"
)

type loginUsecase struct {
	repo entity.UserInterface
}

func New(repo entity.UserInterface) *loginUsecase {
	return &loginUsecase{
		repo: repo,
	}
}

func (login *loginUsecase) Login(userData entity.Login) (token string, err error) {
	result, err := login.repo.SelectUserByEmail(userData.Email)

	if err != nil {
		return token, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(userData.Password))
	if err != nil {
		return token, err
	}

	return middlewares.CreateToken(int(result.Id), result.Role)
}
