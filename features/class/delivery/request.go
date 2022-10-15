package delivery

import "immersiveProject/features/class/entity"

type ClassRequest struct {
	Name        string `json:"name" form:"name"`
	JumlahKelas string `json:"jumlah" form:"jumlah"`
	MulaiKelas  string `json:"mulai" form:"mulai"`
	AkhirKelas  string `json:"akhir" form:"akhir"`
}

func RequestToEntity(request ClassRequest) entity.ClassEntity {
	return entity.ClassEntity{
		Name:        request.Name,
		JumlahKelas: request.JumlahKelas,
		MulaiKelas:  request.MulaiKelas,
		AkhirKelas:  request.AkhirKelas,
	}
}
