package usecase

import (
	"errors"
	"immersiveProject/features/teams"
)

type bookService struct {
	dataBook teams.DataInterface
}

func New(data teams.DataInterface) teams.ServiceInterface {
	return &bookService{
		dataBook: data,
	}

}

func (service *bookService) GetAll() ([]teams.TeamCore, error) {

	dataAll, err := service.dataBook.SelectAll()
	if err != nil {
		return nil, errors.New("failed get all data")
	} else if len(dataAll) == 0 {
		return nil, errors.New("data is still empty")
	} else {
		return dataAll, nil
	}

}

func (service *bookService) GetById(param int) (teams.TeamCore, error) {

	dataId, err := service.dataBook.SelectById(param)
	if err != nil {
		return teams.TeamCore{}, err
	}

	return dataId, nil

}

func (service *bookService) PostData(data teams.TeamCore, token int) (int, error) {

	//&& data.Author != "" && data.Publisher != "" && data.Page != 0 {
	if data.Name != "" {
		add, err := service.dataBook.CreateData(data, token)
		if err != nil || add == 0 {
			return -1, err
		} else {
			return 1, nil
		}
	} else {
		return -1, errors.New("all input data must be filled")
	}

}

func (service *bookService) PutData(param, token int, data teams.TeamCore) (int, error) {

	row, err := service.dataBook.UpdateData(param, token, data)
	if err != nil || row == 0 {
		return -1, err
	}

	return 1, nil

}

func (service *bookService) DeleteData(param, token int) (int, error) {

	row, err := service.dataBook.DelData(param, token)
	if err != nil || row == 0 {
		return -1, err
	}

	return 1, nil

}

// type TeamUsecase struct {
// 	Repo teams.UsecaseTeam
// }

// func New(repo teams.UsecaseTeam) *TeamUsecase {
// 	return &TeamUsecase{
// 		Repo: repo,
// 	}
// }

// func (usecase *TeamUsecase) Create(data teams.Core) (err error) {
// 	err = usecase.Repo.Insert(data)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
