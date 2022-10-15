package entity

type Log struct {
	LogID             int
	Feedback          string
	StatusLog         string
	File              string
	MenteeName        string
	Class             string
	EducationMajor    string
	EducationGraduate string
	Email             string
	Telegram          string
	Phone             string
}

type Mentee struct {
	ID                uint
	FullName          string
	ClassID           uint
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
	Log               []Log
}

type UsecaseLog interface {
	InsertLog(logInsert Log) (row int, err error)
	GetLog() (logInsert []Log, err error)
}

type InterfaceLog interface {
	FindLog() (logInsert []Log, err error)
	CreateLog(logInsert Log) (row int, err error)
}
