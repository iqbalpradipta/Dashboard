package delivery

import (
	"immersiveProject/config"
	"immersiveProject/features/log/entity"
	"immersiveProject/middlewares"
	"immersiveProject/utils/helper"

	// "log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type loghandler struct {
	LogInterface entity.InterfaceLog
}

func New(log entity.InterfaceLog) *loghandler {
	return &loghandler{
		LogInterface: log,
	}
}

func (handler *loghandler) FindLog(c echo.Context) error {
	result, err := handler.LogInterface.FindLog()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("internal server error"))
	}
	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("succses get Log", CoreList(result)))
}

func (handler *loghandler) Createlog(c echo.Context) error {
	logToken, errToken := middlewares.ExtractToken(c)
	if logToken == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed insert data"))
	}

	logs := LogRequest{}
	errBind := c.Bind(&logs)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error bind log"))
	}

	fileData, fileInfo, fileErr := c.Request().FormFile("file")
	if fileErr == http.ErrMissingFile || fileErr != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed to get file"))
	}

	fileExtension, errFileExtention := helper.CheckfileExtension(fileInfo.Filename, config.ContentDocuments)
	if errFileExtention != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("file extension error"))
	}

	errFileSize := helper.CheckFileSize(fileInfo.Size, config.ContentDocuments)
	if errFileSize != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("file size error"))
	}

	filename := strconv.FormatInt(time.Now().Unix(), 10)+ fileExtension
	file, errUploadFile := helper.UploadPDFToS3(config.ContentDocuments, filename, config.ContentDocuments, fileData)

	if errUploadFile != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed to upload"))
	}

	logsCore := ToCoreRequest(logs)
	logsCore.LogID = logToken
	logsCore.File = file

	_, err := handler.LogInterface.CreateLog(logsCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed insert logs"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("insert logs succses"))
}
