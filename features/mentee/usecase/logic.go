package usecase

import (
	"immersiveProject/features/mentee"
)

type menteeUsecase struct {
	menteeData mentee.DataInterface
}

func New(data mentee.DataInterface) mentee.UsecaseInterface {
	return &menteeUsecase{
		menteeData: data,
	}
}

func (usecase *menteeUsecase) PostMentee(data mentee.Core) (int, error) {
	row, err := usecase.menteeData.AddMentee(data)
	if err != nil {
		return -1, err
	}
	return row, nil
}

func (usecase *menteeUsecase) PutMentee(id int, newData mentee.Core) (int, error) {
	row, err := usecase.menteeData.UpdateMentee(id, newData)
	if err != nil {
		return -1, err
	}
	return row, nil
}

func (usecase *menteeUsecase) GetMentee(class int, status string, category string) ([]mentee.Core, error) {
	data, err := usecase.menteeData.SelectMentee(class, status, category)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (usecase *menteeUsecase) DeleteMentee(id int) (int, error) {
	data, err := usecase.menteeData.DeleteData(id)
	if err != nil {
		return -1, err
	}
	return data, nil
}
