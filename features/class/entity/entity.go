package entity

type ClassEntity struct {
	ClassID     uint
	UserID      uint
	Name        string
	JumlahKelas string
	MulaiKelas  string
	AkhirKelas  string
}

type UsecaseClass interface {
	Create(class ClassEntity) (err error)
	UpdateClass(id int, class ClassEntity) (row int, err error)
	DeleteClass(id int) (row int, err error)
	GetClass() (result []ClassEntity, err error)
}

type RepoClass interface {
	Insert(class ClassEntity) (affectedRow int, err error)
	UpdateData(id int, class ClassEntity) (row int, err error)
	DeleteData(id int) (row int, err error)
	FindAll() (result []ClassEntity, err error)
}
