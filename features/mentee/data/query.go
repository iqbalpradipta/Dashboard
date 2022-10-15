package data

import (
	"immersiveProject/features/mentee"

	"gorm.io/gorm"
)

type menteeData struct {
	db *gorm.DB
}

func New(db *gorm.DB) mentee.DataInterface {
	return &menteeData{
		db: db,
	}
}

func (repo *menteeData) AddMentee(data mentee.Core) (int, error) {
	dataModel := fromCore(data)
	tx := repo.db.Create(&dataModel)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *menteeData) UpdateMentee(id int, newData mentee.Core) (int, error) {
	dataModel := fromCore(newData)
	tx := repo.db.Model(&Mentee{}).Where("id = ?", id).Updates(dataModel)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return 1, nil
}

func (repo *menteeData) SelectMentee(class int, status string, category string) ([]mentee.Core, error) {
	//tx := repo.db.Where("class_id = ? OR status_id = ? OR category = ?", classID, statusID, category).Find(&dataMentee)
	var dataMentee []Mentee
	if category != "" {
		var dataByCate []Mentee
		txCate := repo.db.Where("education_category = ?", category).Preload("Class").Find(&dataByCate)
		if txCate.Error != nil {
			return nil, txCate.Error
		}
		return toCoreList(dataByCate), nil

	} else if status != "" {
		var dataByStatus []Mentee
		txState := repo.db.Where("status = ?", status).Preload("Class").Find(&dataByStatus)
		if txState.Error != nil {
			return nil, txState.Error
		}
		return toCoreList(dataByStatus), nil

	} else if class != 0 {
		var dataByClass []Mentee
		txClass := repo.db.Where("class_id = ?", class).Preload("Class").Find(&dataByClass)
		if txClass.Error != nil {
			return nil, txClass.Error
		}
		return toCoreList(dataByClass), nil

	} else {
		txAll := repo.db.Preload("Class").Find(&dataMentee)
		if txAll.Error != nil {
			return nil, txAll.Error
		}
		return toCoreList(dataMentee), nil
	}

}

func (repo *menteeData) DeleteData(id int) (int, error) {
	var deleteData Mentee
	tx := repo.db.Where("id = ?", id).Delete(&deleteData)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return 1, nil
}
