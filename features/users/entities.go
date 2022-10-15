package users

import (
	"time"
)

type UserCore struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	Status    string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Team      []TeamCore
}

type TeamCore struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UserID    uint
	UpdatedAt time.Time
	DeletedAt time.Time
}

type ServiceInterface interface {
	GetAll(token int) (data []UserCore, err error)
	GetById(param, token int) (data UserCore, err error)
	PostData(data UserCore) (int, error)
	PutData(param, token int, data UserCore) (int, error)
	DeleteData(id int) (int, error)
}

type DataInterface interface {
	SelectAll(token int) (data []UserCore, err error)
	SelectById(param, token int) (data UserCore, err error)
	CreateData(data UserCore) (int, error)
	UpdateData(param int, data UserCore) (int, error)
	DelData(id int) (int, error)
}
