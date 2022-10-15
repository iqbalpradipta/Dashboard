package data

import (
	entity "immersiveProject/features/login/entity"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Role     string
	Password string
}

func ToCore(userModel User) entity.Login {
	return entity.Login{
		Id:       userModel.ID,
		Role:     userModel.Role,
		Email:    userModel.Email,
		Password: userModel.Password,
	}
}
