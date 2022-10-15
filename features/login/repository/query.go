package repository

import (
	"gorm.io/gorm"
	entity "immersiveProject/features/login/entity"
	model "immersiveProject/features/login/data"
)
type loginRepo struct{
	db *gorm.DB
}

func New(db *gorm.DB) *loginRepo {
	return &loginRepo{
		db: db,
	}
}

func (repo *loginRepo) SelectUserByEmail(email string) (login entity.Login, err error) {
	userModel := model.User{}
	tx := repo.db.Where("email = ?", email).First(&userModel)
	err = tx.Error

	if err != nil {
		return login, err
	}

	login = model.ToCore(userModel)
	return login, err
}

