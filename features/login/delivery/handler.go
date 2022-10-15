package delivery

import (
	entity "immersiveProject/features/login/entity"
	"immersiveProject/utils/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)
type loginHandler struct{
	Usecase entity.UsecaseLogin
}

func New(usecase entity.UsecaseLogin) *loginHandler {
	return &loginHandler{
		Usecase: usecase,
	}
}

func (repo *loginHandler) Login(c echo.Context)error {
	loginRequest := loginRequest{}
	errBind := c.Bind(&loginRequest)

	if errBind != nil{
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper(errBind.Error()))
	}
	loginEntity := ToEntity(loginRequest)
	token, err := repo.Usecase.Login(loginEntity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("Login succses", token))
}