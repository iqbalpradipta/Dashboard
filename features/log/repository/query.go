package repository

import (
	"errors"
	"immersiveProject/features/log/data"
	"immersiveProject/features/log/entity"

	"gorm.io/gorm"
)

type LogRepo struct{
	db *gorm.DB
}

func New(db *gorm.DB) entity.InterfaceLog {
	return &LogRepo {
		db: db,
	}
}

func (repo *LogRepo) FindLog() ([]entity.Log, error) {
	var logModels  []data.Log
	tx := repo.db.Find(&logModels)
	if tx.Error != nil {
		return nil, tx.Error
	}
	Logs := data.CoreList(logModels)
	return Logs, nil
}

func (repo *LogRepo) CreateLog(logCreate entity.Log)  (int, error){
	LogModel := data.FromCore(logCreate)
	tx := repo.db.Create(&LogModel)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("failed to create log")
	}
	return int(tx.RowsAffected), nil
}