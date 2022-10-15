package data

import (
	"errors"
	"immersiveProject/features/teams"

	// _ "immersiveProject/features/teams/data"

	"gorm.io/gorm"
)

type teamData struct {
	DB *gorm.DB
}

func New(conn *gorm.DB) teams.DataInterface {
	return &teamData{
		DB: conn,
	}
}

func (repo *teamData) SelectAll() ([]teams.TeamCore, error) {

	var dataTeam []Team
	tx := repo.DB.Find(&dataTeam)
	if tx.Error != nil {
		return nil, tx.Error
	}

	dataTeamCore := toTeamList(dataTeam)

	return dataTeamCore, nil

}

func (repo *teamData) SelectById(param int) (teams.TeamCore, error) {

	var data Team
	tx := repo.DB.First(&data, param)
	if tx.Error != nil {
		return teams.TeamCore{}, tx.Error
	}

	teamId := data.ToTeamUser()
	return teamId, nil

}

func (repo *teamData) CreateData(data teams.TeamCore, token int) (int, error) {

	var datacheck User
	txcheck := repo.DB.Where("ID=?", token).First(&datacheck)
	if txcheck.Error != nil {
		return -1, errors.New("error tx")
	}

	if int(datacheck.ID) != token {
		return -1, errors.New("not have access")
	}

	dataModel := InsTeam(data)
	dataModel.UserID = uint(token)
	tx := repo.DB.Create(&dataModel)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil

}

func (repo *teamData) UpdateData(param, token int, dataUpdate teams.TeamCore) (int, error) {

	var datacheck Team
	tx := repo.DB.First(&datacheck, param)
	if tx.Error != nil {
		return -1, tx.Error
	}

	TeamId := datacheck.ToTeamUser()

	if TeamId.UserID == uint(token) {
		var data Team
		data.Name = dataUpdate.Name

		var team Team
		team.ID = dataUpdate.ID
		txUpdateId := repo.DB.Model(&TeamId).Updates(data)
		if txUpdateId.Error != nil {
			return -1, txUpdateId.Error
		}

		var err error

		return int(txUpdateId.RowsAffected), err
	} else {
		return -1, errors.New("not have access")
	}

}

func (repo *teamData) DelData(param, token int) (int, error) {

	var datacheck Team
	tx := repo.DB.First(&datacheck, param)
	if tx.Error != nil {
		return -1, tx.Error
	}

	teamId := datacheck.ToTeamUser()

	if teamId.UserID == uint(token) {
		var data Team
		txDelId := repo.DB.Delete(&data, param)
		if txDelId.Error != nil {
			return -1, txDelId.Error
		}

		var err error

		return int(txDelId.RowsAffected), err
	} else {
		return -1, errors.New("not access")
	}

}
