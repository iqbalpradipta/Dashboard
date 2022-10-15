package repository

import (
	"errors"
	"immersiveProject/features/class/data"
	"immersiveProject/features/class/entity"

	"gorm.io/gorm"
)

type classRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *classRepo {
	return &classRepo{
		db: db,
	}
}

func (repo *classRepo) Insert(class entity.ClassEntity) (affectedRow int, err error) {
	entityModel := data.EntityToModel(class)

	tx := repo.db.Model(&data.Class{}).Create(&entityModel)

	if tx.Error != nil {
		return -1, tx.Error
	}

	if tx.RowsAffected < 1 {
		return int(tx.RowsAffected), errors.New("failed insert class")
	}

	return int(tx.RowsAffected), nil
}

func (repo *classRepo) DeleteData(id int) (row int, err error) {

	classModel := data.Class{}

	tx := repo.db.Where("id = ?", id).Delete(&classModel)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("failed  delete class")
	}
	return int(tx.RowsAffected), nil
}

func (repo *classRepo) UpdateData(id int, class entity.ClassEntity) (row int, err error) {
	classModel := data.EntityToModel(class)
	tx := repo.db.Model(&data.Class{}).Where("id = ?", id).Updates(&classModel)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("failed to update data")
	}
	return 1, nil
}

func (repo *classRepo) FindAll() (result []entity.ClassEntity, err error) {
	classModels := []data.Class{}
	tx := repo.db.Find(&classModels)
	if tx.Error != nil {
		return result, tx.Error
	}

	for _, class := range classModels {
		result = append(result, data.ModelToEntity(class))
	}

	return result, nil
}
