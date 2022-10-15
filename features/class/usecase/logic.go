package usecase

import (
	"immersiveProject/features/class/entity"
)

type classUsecase struct {
	Repo entity.RepoClass
}

func New(repo entity.RepoClass) *classUsecase {
	return &classUsecase{
		Repo: repo,
	}
}

func (usecase *classUsecase) Create(class entity.ClassEntity) (err error) {
	_, err = usecase.Repo.Insert(class)

	if err != nil {
		return err
	}

	return nil
}

func (usecase *classUsecase) UpdateClass(id int, class entity.ClassEntity) (row int, err error) {
	classMap := make(map[string]interface{})
	if class.Name != "" {
		classMap["name"] = &class.Name
	}
	if class.JumlahKelas != "" {
		classMap["jumlah"] = &class.JumlahKelas
	}
	if class.MulaiKelas != "" {
		classMap["mulai"] = &class.MulaiKelas
	}
	if class.AkhirKelas != "" {
		classMap["akhir"] = &class.AkhirKelas
	}

	result, err := usecase.Repo.UpdateData(id, class)
	if err != nil {
		return -1, err
	}
	return result, nil
}

func (usecase *classUsecase) DeleteClass(id int) (row int, err error) {
	result, err := usecase.Repo.DeleteData(id)
	if err != nil {
		return -1, err
	}
	return result, nil
}

func (usecase *classUsecase) GetClass() ([]entity.ClassEntity, error) {
	result, err := usecase.Repo.FindAll()

	return result, err
}
