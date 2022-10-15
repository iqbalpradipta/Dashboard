package delivery

import (
	entity "immersiveProject/features/login/entity"
)
type loginRequest struct {
	Email		string	`json:"email" form:"email"`
	Password	string	`json:"password" form:"password"`
}

func ToEntity(request loginRequest) entity.Login {
	return entity.Login{
		Email: request.Email,
		Password: request.Password,
	}
}