package data

import (
	"immersiveProject/features/class/entity"

	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	UserID      uint
	Name        string
	JumlahKelas string
	MulaiKelas  string
	AkhirKelas  string
	User        User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Mentee      []Mentee
}

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

type Mentee struct {
	gorm.Model
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
	Class             Class
}

func ModelToEntity(classModel Class) entity.ClassEntity {
	return entity.ClassEntity{
		ClassID:     classModel.ID,
		UserID:      classModel.UserID,
		Name:        classModel.Name,
		JumlahKelas: classModel.JumlahKelas,
		MulaiKelas:  classModel.MulaiKelas,
		AkhirKelas:  classModel.AkhirKelas,
	}
}

func EntityToModel(classEntity entity.ClassEntity) Class {
	return Class{
		UserID:      classEntity.UserID,
		Name:        classEntity.Name,
		JumlahKelas: classEntity.JumlahKelas,
		MulaiKelas:  classEntity.MulaiKelas,
		AkhirKelas:  classEntity.AkhirKelas,
	}
}
