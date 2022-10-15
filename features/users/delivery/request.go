package delivery

import "immersiveProject/features/users"

type UserReq struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Status   string `json:"status" form:"status"`
	Role     string `json:"role" form:"role"`
	Team     string `json:"team" form:"team"`
}

func ToCore(req UserReq) users.UserCore {

	return users.UserCore{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Status:   req.Status,
		Role:     req.Role,
	}

}
