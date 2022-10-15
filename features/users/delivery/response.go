package delivery

import "immersiveProject/features/users"

type Respon struct {
	ID     uint   `json:"id" form:"id"`
	Name   string `json:"name" form:"name"`
	Email  string `json:"email" form:"email"`
	Status string `json:"status" form:"status"`
	Role   string `json:"role" form:"role"`
	Team   []users.TeamCore
}

func toResponId(data users.UserCore) Respon {

	return Respon{
		ID:     data.ID,
		Name:   data.Name,
		Email:  data.Email,
		Status: data.Status,
		Role:   data.Role,
		Team:   data.Team,
	}

}

func toResponList(data []users.UserCore) []Respon {

	var respon []Respon
	for _, v := range data {
		respon = append(respon, Respon{
			ID:     v.ID,
			Name:   v.Name,
			Email:  v.Email,
			Status: v.Status,
			Role:   v.Role,
			Team:   v.Team,
		})
	}

	return respon
}
