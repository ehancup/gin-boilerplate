package src

import (
	// "fmt"
	"gin-boilerplate/config"
	"gin-boilerplate/src/utils/logger"

	// "time"

	log "github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	

	_ "gin-boilerplate/docs"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
)

func BoostrapApp() {
	// testing

	err := godotenv.Load()

	if err != nil {
		logger.Fatal("Error loading .env file")
	}

	appCfg := config.GetConfig().App
	if appCfg.Mode == "debug" {
        logger.Log.SetLevel(log.DebugLevel)
		logger.Debug("debug mode")
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	/*
		user commmented code below to use database
		dont forget to import "gin-boilerplate/src/database" and "flag"
	*/

	// database.ConnectDatabase()

	// noDbs := flag.Bool("no-migrate", false, "set to not migrate database.")
	// flag.Parse()

	// if !*noDbs {
	// 	database.Migrate()
	// } else {
	// 	logger.Info("database set to not migrate")

	// }
	
	app := gin.New()

	app.Use(gin.Recovery())

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	InitRoutes(app)

	logger.Info("Check documentation ->", "url", appCfg.Url + "/swagger/index.html")

	if err := app.Run(appCfg.Port); err != nil {
		logger.Fatal("[ERR] fail starting servr :", "err", err)
	}
}
