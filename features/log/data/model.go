package data

import (
	"immersiveProject/features/log/entity"

	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	UserID    uint
	MenteeID  uint
	Feedback  string
	StatusLog string
	File      string
	Mentee    Mentee
}

type Mentee struct {
	gorm.Model
	UserID            uint
	MenteeID          uint
	Status            string
	Address           string
	HomeAddress       string
	Email             string `gorm:"unique"`
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

func FromCore(logCore entity.Log) Log {
	logModel := Log{
		Feedback:  logCore.Feedback,
		StatusLog: logCore.StatusLog,
		File:      logCore.File,
	}
	return logModel
}

func (logCore *Log) ToCore() entity.Log {
	return entity.Log{
		LogID:     int(logCore.ID),
		Feedback:  logCore.Feedback,
		StatusLog: logCore.StatusLog,
		File:      logCore.File,
	}
}

func CoreList(logCore []Log) []entity.Log {
	var LogCore []entity.Log
	for key := range logCore {
		LogCore = append(LogCore, logCore[key].ToCore())
	}
	return LogCore
}
