package delivery

import (
	"immersiveProject/features/teams"
	"immersiveProject/middlewares"
	"immersiveProject/utils/helper"

	"github.com/labstack/echo/v4"
)

type TeamHandler struct {
	data teams.ServiceInterface
}

func New(e *echo.Echo, usecase teams.ServiceInterface) {

	handler := TeamHandler{
		data: usecase,
	}

	e.GET("/teams", handler.GetAllWithJWT, middlewares.JWTMiddleware())
	e.GET("/teams/:id", handler.GetByIdWithJWT, middlewares.JWTMiddleware())
	e.POST("/teams", handler.AddUser, middlewares.JWTMiddleware())
	e.PUT("/teams/:id", handler.PutDataWithJWT, middlewares.JWTMiddleware())
	e.DELETE("/teams/:id", handler.DeldateWithJWT, middlewares.JWTMiddleware())

}

func (th *TeamHandler) GetAllWithJWT(e echo.Context) error {

	res, err := th.data.GetAll()
	if err != nil {
		return e.JSON(400, helper.FailedResponseHelper("tidak ada data"))
	}

	respon := toResponList(res)

	return e.JSON(200, helper.SuccessDataResponseHelper("succes get all data", respon))

}

func (th *TeamHandler) GetByIdWithJWT(e echo.Context) error {

	id := helper.ParamInt(e)
	if id == -1 {
		return e.JSON(400, helper.FailedResponseHelper("param must be number"))
	}

	res, err := th.data.GetById(id)
	if err != nil {
		return e.JSON(400, helper.FailedResponseHelper("id not found"))
	}

	respon := toRespon(res)

	return e.JSON(200, helper.SuccessDataResponseHelper("succes get data by id", respon))

}

func (th *TeamHandler) AddUser(e echo.Context) error {

	idToken, _ := middlewares.ExtractToken(e)
	var req Request
	err := e.Bind(&req)
	if err != nil {
		return e.JSON(400, helper.FailedResponseHelper("err"))
	}

	add := toCore(req)
	row, _ := th.data.PostData(add, idToken)
	if row == 1 {
		return e.JSON(200, helper.SuccessResponseHelper("succes insert data"))
	} else {
		return e.JSON(400, helper.FailedResponseHelper("failed insert teams"))
	}

}

func (th *TeamHandler) PutDataWithJWT(e echo.Context) error {

	idToken, _ := middlewares.ExtractToken(e)
	id := helper.ParamInt(e)
	if id == -1 {
		return e.JSON(400, helper.FailedResponseHelper("id not found"))
	}

	var req Request
	err := e.Bind(&req)
	if err != nil {
		return e.JSON(400, helper.FailedResponseHelper("error bind"))
	}

	var add teams.TeamCore
	if req.Name != "" {
		add.Name = req.Name
	}

	add.ID = uint(id)

	row, _ := th.data.PutData(id, idToken, add)
	if row == 1 {
		return e.JSON(200, helper.SuccessResponseHelper("succes update data"))
	} else {
		return e.JSON(400, helper.FailedResponseHelper("not have access"))
	}

}

func (th *TeamHandler) DeldateWithJWT(e echo.Context) error {

	idToken, _ := middlewares.ExtractToken(e)
	id := helper.ParamInt(e)
	if id == -1 {
		return e.JSON(400, helper.FailedResponseHelper("param must be number"))
	}

	row, _ := th.data.DeleteData(id, idToken)
	if row == 1 {
		return e.JSON(200, helper.SuccessResponseHelper("succes delete data"))
	} else {
		return e.JSON(400, helper.FailedResponseHelper("not have access"))
	}

}
