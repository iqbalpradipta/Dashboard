package teams

import "time"

type TeamCore struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UserID    uint
	UpdatedAt time.Time
	DeletedAt time.Time
}

type DataInterface interface {
	SelectAll() (data []TeamCore, err error)
	SelectById(param int) (data TeamCore, err error)
	CreateData(data TeamCore, token int) (int, error)
	UpdateData(param, token int, data TeamCore) (int, error)
	DelData(param, token int) (int, error)
}

type ServiceInterface interface {
	GetAll() (data []TeamCore, err error)
	GetById(param int) (data TeamCore, err error)
	PostData(data TeamCore, token int) (int, error)
	PutData(param, token int, data TeamCore) (int, error)
	DeleteData(param, token int) (int, error)
}
