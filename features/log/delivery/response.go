package delivery

import "immersiveProject/features/log/entity"

type LogResponse struct {
	Feedback          string `json:"feedback"`
	StatusLog         string `json:"statuslog"`
	File              string `json:"file"`
	MenteeName        string
	Class             string
	EducationMajor    string
	EducationGraduate string
	Email             string
	Telegram          string
	Phone             string
}

func FromCore(logData entity.Log) LogResponse {
	return LogResponse{
		Feedback:          logData.Feedback,
		StatusLog:         logData.StatusLog,
		File:              logData.File,
		MenteeName:        logData.MenteeName,
		Class:             logData.Class,
		EducationMajor:    logData.EducationMajor,
		EducationGraduate: logData.EducationGraduate,
		Email:             logData.Email,
		Telegram:          logData.Telegram,
		Phone:             logData.Phone,
	}
}

func CoreList(logData []entity.Log) []LogResponse {
	var ResponseLog []LogResponse
	for _, v := range logData {
		ResponseLog = append(ResponseLog, FromCore(v))
	}
	return ResponseLog
}

func CoreResponse(logData entity.Log) LogResponse {
	ResponseLog := LogResponse{
		Feedback:          logData.Feedback,
		StatusLog:         logData.StatusLog,
		MenteeName:        logData.MenteeName,
		Class:             logData.Class,
		EducationMajor:    logData.EducationMajor,
		EducationGraduate: logData.EducationGraduate,
		Email:             logData.Email,
		Telegram:          logData.Telegram,
		Phone:             logData.Phone,
	}
	return ResponseLog
}
