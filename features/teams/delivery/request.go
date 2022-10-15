package delivery

import "immersiveProject/features/teams"

type Request struct {
	Name string `json:"name" form:"name"`
}

func toCore(req Request) teams.TeamCore {

	var res = teams.TeamCore{
		Name: req.Name,
	}

	return res

}
