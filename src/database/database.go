package database

import (
	// "gin-gorm/app/models"
	"gin-boilerplate/config"
	"gin-boilerplate/src/database/dao"

	// "gin-boilerplate/src/app/users"
	"gin-boilerplate/src/utils/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	var errConnection error
	dbCfg := config.GetConfig().DB

	DB, errConnection = gorm.Open(mysql.Open(dbCfg.DSN), &gorm.Config{})

	if errConnection != nil {
		logger.Fatal("Cant connect to database")
	}

	logger.Info("Success connected to database", "driver")

}

func Migrate() {
	err := DB.AutoMigrate(
		&dao.UserEntity{},
	)

	if err != nil {
		logger.Fatal("Migration Failed", "err", err.Error())
	}
	logger.Info("Database migrated successfully.")
}
