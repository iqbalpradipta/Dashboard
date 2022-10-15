package data

import (
	"immersiveProject/features/teams"
	"time"

	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	UserID uint
	Name   string `json:"name" form:"name"`
}

type User struct {
	Email     string `gorm:"primary key"`
	ID        uint   `gorm:"autoIncrement"`
	Name      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func InsTeam(data teams.TeamCore) Team {
	userData := Team{
		Name: data.Name,
	}

	return userData
}

func (dataTeam *Team) ToTeamUser() teams.TeamCore {

	dataTeamCore := teams.TeamCore{
		ID:        dataTeam.ID,
		UserID:    dataTeam.UserID,
		Name:      dataTeam.Name,
		CreatedAt: dataTeam.CreatedAt,
		UpdatedAt: dataTeam.UpdatedAt,
		DeletedAt: dataTeam.DeletedAt.Time,
	}

	return dataTeamCore

}

func toTeamList(data []Team) []teams.TeamCore {
	var dataCore []teams.TeamCore
	for i := 0; i < len(data); i++ {
		dataCore = append(dataCore, data[i].ToTeamUser())
	}
	return dataCore
}
