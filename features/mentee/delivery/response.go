package delivery

import (
	"immersiveProject/features/mentee"
)

type MenteeResponse struct {
	ID                uint   `json:"id"`
	FullName          string `json:"fullname"`
	ClassName         string `json:"classname"`
	Status            string `json:"status"`
	Address           string `json:"address,omitempty"`
	HomeAddress       string `json:"homeaddress,omitempty"`
	Email             string `json:"email,omitempty"`
	Gender            string `json:"gender,omitempty"`
	Telegram          string `json:"telegram,omitempty"`
	Phone             string `json:"phone,omitempty"`
	EmergencyName     string `json:"emergencyname,omitempty"`
	EmergencyPhone    string `json:"emergencyphone,omitempty"`
	EmergencyStatus   string `json:"emergencystatus,omitempty"`
	EducationCategory string `json:"educationcategory,omitempty"`
	EducationMajor    string `json:"educationmajor,omitempty"`
	EducationGraduate string `json:"educationgraduate,omitempty"`
}

func fromCore(data mentee.Core) MenteeResponse {
	return MenteeResponse{
		ID:                data.ID,
		FullName:          data.FullName,
		ClassName:         data.ClassName,
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

func fromCoreList(data []mentee.Core) []MenteeResponse {
	var dataRes []MenteeResponse
	for _, v := range data {
		dataRes = append(dataRes, MenteeResponse{
			ID:                v.ID,
			FullName:          v.FullName,
			ClassName:         v.ClassName,
			Status:            v.Status,
			Address:           v.Address,
			HomeAddress:       v.HomeAddress,
			Email:             v.Email,
			Gender:            v.Gender,
			Telegram:          v.Telegram,
			Phone:             v.Phone,
			EmergencyName:     v.EmergencyName,
			EmergencyPhone:    v.EmergencyPhone,
			EmergencyStatus:   v.EmergencyStatus,
			EducationCategory: v.EducationCategory,
			EducationMajor:    v.EducationMajor,
			EducationGraduate: v.EducationGraduate,
		})
	}

	return dataRes
}
