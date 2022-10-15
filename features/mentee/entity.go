package mentee

type Core struct {
	ID                uint
	FullName          string
	ClassID           uint
	ClassName         string
	Status            string
	Address           string
	HomeAddress       string
	Email             string
	Gender            string
	Telegram          string
	Phone             string
	EmergencyName     string
	EmergencyPhone    string
	EmergencyStatus   string
	EducationCategory string
	EducationMajor    string
	EducationGraduate string
}

type Class struct {
	ID   uint
	Name string
}

type UsecaseInterface interface {
	PostMentee(data Core) (int, error)
	GetMentee(class int, status string, category string) (data []Core, err error)
	PutMentee(id int, newData Core) (row int, err error)
	DeleteMentee(id int) (int, error)
}

type DataInterface interface {
	AddMentee(data Core) (int, error)
	SelectMentee(class int, status string, category string) (data []Core, err error)
	UpdateMentee(id int, newData Core) (row int, err error)
	DeleteData(id int) (int, error)
}
