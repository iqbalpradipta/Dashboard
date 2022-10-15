package usecase

import (
	"errors"
	"immersiveProject/features/log/entity"
)

type logLogic struct{
	logData entity.InterfaceLog
}

func New(logic entity.InterfaceLog) entity.InterfaceLog {
	return &logLogic{
		logData: logic,
	}
}

func (logic *logLogic) CreateLog(logCreate entity.Log) (int, error){
	if logCreate.Feedback == "" || logCreate.StatusLog == "" || logCreate.File == "" {
		return -1, errors.New("log not fuel field")
	}
	
	result, err := logic.logData.CreateLog(logCreate)
	if err != nil{
		return 0, errors.New("failed to insert data")
	}
	return result, nil
}

func (logic *logLogic) FindLog() ([]entity.Log, error) {
	result, err := logic.logData.FindLog()
	return result, err
}