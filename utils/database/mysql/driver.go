package mysql

import (
	"fmt"
	"immersiveProject/config"
	classModel "immersiveProject/features/class/data"
	menteeModel "immersiveProject/features/mentee/data"
	_teamsModel "immersiveProject/features/teams/data"
	_userModel "immersiveProject/features/users/data"
	_logModel "immersiveProject/features/log/data"

	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("error connect to DB", err)
	}

	autoMigrate(db)

	return db
}

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		new(_teamsModel.Team),
		new(_userModel.User),
		new(classModel.Class),
		new(menteeModel.Mentee),
		new(_logModel.Log),
		
	)
}
