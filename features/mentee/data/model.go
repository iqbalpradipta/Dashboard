package data

import (
	"immersiveProject/features/mentee"

	"gorm.io/gorm"
)

type Mentee struct {
	gorm.Model
	FullName          string
	ClassID           uint
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
	Class             Class
}

type Class struct {
	gorm.Model
	Name    string
	Mentees []Mentee
}

func fromCore(data mentee.Core) Mentee {
	dataModel := Mentee{
		FullName:          data.FullName,
		ClassID:           data.ClassID,
		Status:            data.Status,
		Address:           data.Address,
		HomeAddress:       data.HomeAddress,
		Email:             data.Email,
		Gender:            data.Gender,
		Telegram:          data.Telegram,
		Phone:             data.Phone,
		EmergencyName:     data.EmergencyName,
		EmergencyPhone:    data.EmergencyPhone,
		EmergencyStatus:   data.EmergencyStatus,
		EducationCategory: data.EducationCategory,
		EducationMajor:    data.EducationMajor,
		EducationGraduate: data.EducationGraduate,
	}
	return dataModel
}

func (data *Mentee) toCore() mentee.Core {
	return mentee.Core{
		ID:                data.ID,
		FullName:          data.FullName,
		ClassID:           data.ClassID,
		ClassName:         data.Class.Name,
		Status:            data.Status,
		Address:           data.Address,
		HomeAddress:       data.HomeAddress,
		Email:             data.Email,
		Gender:            data.Gender,
		Telegram:          data.Telegram,
		Phone:             data.Phone,
		EmergencyName:     data.EmergencyName,
		EmergencyPhone:    data.EmergencyPhone,
		EmergencyStatus:   data.EmergencyStatus,
		EducationCategory: data.EducationCategory,
		EducationMajor:    data.EducationMajor,
		EducationGraduate: data.EducationGraduate,
	}
}

func toCoreList(data []Mentee) []mentee.Core {
	var dataCore []mentee.Core
	for key := range data {
		dataCore = append(dataCore, data[key].toCore())
	}
	return dataCore
}
