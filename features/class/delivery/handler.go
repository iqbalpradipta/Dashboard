package delivery

import (
	"immersiveProject/features/class/entity"
	"immersiveProject/middlewares"
	"immersiveProject/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type classhandler struct {
	Usecase entity.UsecaseClass
}

func New(usecase entity.UsecaseClass) *classhandler {
	return &classhandler{
		Usecase: usecase,
	}
}

func (repo *classhandler) Create(c echo.Context) error {
	var userRequest ClassRequest

	id, err := middlewares.ExtractToken(c)
	if err != nil {
		return c.JSON(http.StatusForbidden, helper.FailedResponseHelper(err.Error()))
	}

	err = c.Bind(&userRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper(err.Error()))
	}
	userEntity := RequestToEntity(userRequest)
	userEntity.UserID = uint(id)

	err = repo.Usecase.Create(userEntity)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("Succses Create Class"))
}

func (repo *classhandler) UpdateClass(c echo.Context) error {

	id := c.Param("id")
	idCnv, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("param must be number"))
	}

	var classUpdate ClassRequest
	errBind := c.Bind(&classUpdate)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind data"))
	}

	result, err := repo.Usecase.UpdateClass(idCnv, RequestToEntity(classUpdate))
	if err != nil {
		return c.JSON(500, helper.FailedResponseHelper("failed update class"))
	}
	return c.JSON(200, helper.SuccessDataResponseHelper("success update class", result))
}

func (repo *classhandler) DeleteClass(c echo.Context) error {

	id := c.Param("id")
	idCnv, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("param must be number"))
	}

	row, err := repo.Usecase.DeleteClass(idCnv)
	if err != nil || row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed delete class"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("success delete class"))
}

func (repo *classhandler) GetClass(c echo.Context) error {
	result, err := repo.Usecase.GetClass()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("failed get class"))
	}

	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("succses get class", EntityList(result)))
}