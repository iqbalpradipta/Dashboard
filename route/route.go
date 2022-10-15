package route

import (
	"immersiveProject/config"
	"immersiveProject/middlewares"


	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	loginhandler "immersiveProject/features/login/delivery"
	loginrepo "immersiveProject/features/login/repository"
	loginusecase "immersiveProject/features/login/usecase"

	userData "immersiveProject/features/users/data"
	userDelivery "immersiveProject/features/users/delivery"
	userUsecase "immersiveProject/features/users/usecase"

	teamData "immersiveProject/features/teams/data"
	teamDelivery "immersiveProject/features/teams/delivery"
	teamUsecase "immersiveProject/features/teams/usecase"

	classhandler "immersiveProject/features/class/delivery"
	classrepo "immersiveProject/features/class/repository"
	classUsecase "immersiveProject/features/class/usecase"

	menteeData "immersiveProject/features/mentee/data"
	menteeDelivery "immersiveProject/features/mentee/delivery"
	menteeUsecase "immersiveProject/features/mentee/usecase"

	loghandler "immersiveProject/features/log/delivery"
	logrepo "immersiveProject/features/log/repository"
	logusecase "immersiveProject/features/log/usecase"
)

func InitRoutes(e *echo.Echo, db *gorm.DB, cfg *config.AppConfig) {
	loginRepo := loginrepo.New(db)
	loginUsecase := loginusecase.New(loginRepo)
	loginHandler := loginhandler.New(loginUsecase)

	userData := userData.New(db)
	userUsecaseFactory := userUsecase.New(userData)
	userDelivery.New(e, userUsecaseFactory)

	teamData := teamData.New(db)
	teamUsecase := teamUsecase.New(teamData)
	teamDelivery.New(e, teamUsecase)

	classrepo := classrepo.New(db)
	classUsecase := classUsecase.New(classrepo)
	classHandler := classhandler.New(classUsecase)

	menteeData := menteeData.New(db)
	menteeUsecaseFactory := menteeUsecase.New(menteeData)
	menteeDelivery.New(e, menteeUsecaseFactory)

	logRepo := logrepo.New(db)
	logUsecase := logusecase.New(logRepo)
	logHandler := loghandler.New(logUsecase)

	/*  Route
	 */

	e.POST("/login", loginHandler.Login)

	e.GET("/class", classHandler.GetClass)
	e.POST("/class", classHandler.Create, middlewares.JWTMiddleware())
	e.PUT("/class/:id", classHandler.UpdateClass, middlewares.JWTMiddleware())
	e.DELETE("/class/:id", classHandler.DeleteClass, middlewares.JWTMiddleware())

	e.GET("/log", logHandler.FindLog, middlewares.JWTMiddleware())
	e.POST("/log", logHandler.Createlog, middlewares.JWTMiddleware())
	

}
