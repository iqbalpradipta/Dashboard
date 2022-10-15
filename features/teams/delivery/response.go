package delivery

import "immersiveProject/features/teams"

type Respon struct {
	Name   string `json:"name" form:"name"`
	UserID uint   `json:"user_id" form:"user_id"`
}

func toRespon(data teams.TeamCore) Respon {

	var res = Respon{
		Name:   data.Name,
		UserID: data.UserID,
	}

	return res

}

func toResponList(data []teams.TeamCore) []Respon {

	var res []Respon
	for _, v := range data {
		res = append(res, toRespon(v))
	}

	return res

}
