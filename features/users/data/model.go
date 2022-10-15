package data

import (
	"immersiveProject/features/users"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Email     string `gorm:"primary key"`
	ID        uint   `gorm:"autoIncrement"`
	Name      string
	Password  string
	Status    string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Teams     []Team `gorm:"foreignkey:UserID;references:ID"`
}

type Team struct {
	gorm.Model
	UserID uint
	Name   string `json:"name" form:"name"`
}

func ModelInsert(data users.UserCore) User {

	userData := User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Status:   data.Status,
		Role:     data.Role,
	}

	return userData

}

func (dataTeamUser *Team) toTeamUser() users.TeamCore {

	dataUser := users.TeamCore{
		ID:        dataTeamUser.ID,
		UserID:    dataTeamUser.UserID,
		Name:      dataTeamUser.Name,
		CreatedAt: dataTeamUser.CreatedAt,
		UpdatedAt: dataTeamUser.UpdatedAt,
		DeletedAt: dataTeamUser.DeletedAt.Time,
	}

	return dataUser

}

func toTeamUserList(data []Team) []users.TeamCore {
	var dataCore []users.TeamCore
	for i := 0; i < len(data); i++ {
		dataCore = append(dataCore, data[i].toTeamUser())
	}
	return dataCore
}

func (data *User) toUserCore(team []Team) users.UserCore {

	dataUser := users.UserCore{
		ID:        data.ID,
		Name:      data.Name,
		Email:     data.Email,
		Password:  data.Password,
		Status:    data.Status,
		Role:      data.Role,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	teamUser := toTeamUserList(team)

	for i, v := range teamUser {
		if v.UserID == dataUser.ID {
			dataUser.Team = append(dataUser.Team, teamUser[i])
		}
	}

	return dataUser
}

func toUserCoreList(data []User, team []Team) []users.UserCore {
	var dataCore []users.UserCore
	for i := 0; i < len(data); i++ {
		dataCore = append(dataCore, data[i].toUserCore(team))
	}
	return dataCore
}
