package delivery

import "immersiveProject/features/class/entity"

type ClassResponse struct{
	ID			uint
	Name		string
	JumlahKelas	string
	MulaiKelas	string
	AkhirKelas	string
}

func EntityToClassResponse(classEntity entity.ClassEntity) ClassResponse {
	return ClassResponse{
		ID: 		classEntity.ClassID,
		Name: 		classEntity.Name,
		JumlahKelas: classEntity.JumlahKelas,
		MulaiKelas: classEntity.MulaiKelas,
		AkhirKelas: classEntity.AkhirKelas,
	}
}

func EntityList(data []entity.ClassEntity) []ClassResponse {
	var Response []ClassResponse
	for _, v := range data {
		Response = append(Response, EntityToClassResponse(v))
	}
	return Response
}