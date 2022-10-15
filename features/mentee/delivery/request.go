package delivery

import "immersiveProject/features/mentee"

type MenteeRequest struct {
	FullName          string `json:"name" form:"name"`
	ClassID           uint   `json:"class_id" form:"class_id"`
	Status            string `json:"status" form:"status"`
	Address           string `json:"address" form:"address"`
	HomeAddress       string `json:"homeaddress" form:"homeaddress"`
	Email             string `json:"email" form:"email"`
	Gender            string `json:"gender" form:"gender"`
	Telegram          string `json:"telegram" form:"telegram"`
	Phone             string `json:"phone" form:"phone"`
	EmergencyName     string `json:"emergencyname" form:"emergencyname"`
	EmergencyPhone    string `json:"emergencyphone" form:"emergencyphone"`
	EmergencyStatus   string `json:"emergencystatus" form:"emergencystatus"`
	EducationCategory string `json:"educationcategory" form:"educationcategory"`
	EducationMajor    string `json:"educationmajor" form:"educationmajor"`
	EducationGraduate string `json:"educationgraduate" form:"educationgraduate"`
}

func toCore(data MenteeRequest) mentee.Core {
	return mentee.Core{
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
}
