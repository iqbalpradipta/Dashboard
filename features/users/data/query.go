package data

import (
	"errors"
	"immersiveProject/features/users"

	"gorm.io/gorm"
)

type userData struct {
	DB *gorm.DB
}

func New(conn *gorm.DB) users.DataInterface {
	return &userData{
		DB: conn,
	}
}

func (repo *userData) team() []Team {

	var dataTeamUser []Team
	tx := repo.DB.Find(&dataTeamUser)
	if tx.Error != nil {
		return nil
	}

	return dataTeamUser

}

func (repo *userData) SelectAll(token int) ([]users.UserCore, error) {

	var datacheck User
	txcheck := repo.DB.Where("ID=?", token).First(&datacheck)
	if txcheck.Error != nil {
		return nil, errors.New("error tx")
	}

	if int(datacheck.ID) != token {
		return nil, errors.New("not have access")
	}

	var dataAll []User
	tx := repo.DB.Find(&dataAll)
	if tx.Error != nil {
		return nil, tx.Error
	}

	teamList := repo.team()

	dataCore := toUserCoreList(dataAll, teamList)
	return dataCore, nil

}

func (repo *userData) SelectById(param, token int) (users.UserCore, error) {

	var datacheck User
	txcheck := repo.DB.Where("ID=?", token).First(&datacheck)
	if txcheck.Error != nil {
		return users.UserCore{}, errors.New("error tx")
	}

	if int(datacheck.ID) != token {
		return users.UserCore{}, errors.New("not have access")
	}

	var data User
	tx := repo.DB.First(&data, param)
	if tx.Error != nil {
		return users.UserCore{}, tx.Error
	}

	teamList := repo.team()

	userId := data.toUserCore(teamList)
	return userId, nil

}

func (repo *userData) CreateData(data users.UserCore) (int, error) {

	dataModel := ModelInsert(data)
	tx := repo.DB.Create(&dataModel)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil

}

func (repo *userData) UpdateData(param int, dataUpdate users.UserCore) (int, error) {

	var data User
	data.Name = dataUpdate.Name
	data.Email = dataUpdate.Email
	data.Password = dataUpdate.Password

	var user User
	user.ID = dataUpdate.ID
	txUpdateId := repo.DB.Model(&user).Where("id = ?", param).Updates(data)
	if txUpdateId.Error != nil {
		return -1, txUpdateId.Error
	}

	var err error

	return int(txUpdateId.RowsAffected), err

}

func (repo *userData) DelData(id int) (int, error) {

	var data User
	txDelId := repo.DB.Where("id = ?", id).Delete(&data)
	if txDelId.Error != nil {
		return -1, txDelId.Error
	}

	return 1, nil

}
