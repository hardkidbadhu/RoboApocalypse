package router

import (
	"github.com/RobotApocalypse/Client"
	"github.com/RobotApocalypse/configuration"
	"github.com/RobotApocalypse/constants"
	"github.com/RobotApocalypse/controllers"
	"github.com/RobotApocalypse/docs"
	"github.com/RobotApocalypse/middleware"
	"github.com/RobotApocalypse/repository"
	"github.com/RobotApocalypse/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(l *logrus.Entry, cfg configuration.Config) *gin.Engine {

	dbConnect := repository.NewDBConnect(cfg, l)
	dbIns := dbConnect.ConnectDB()
	dbConnect.Migrate()

	router := gin.Default()

	newMiddleware := middleware.NewMiddleware(l)

	router.Use(newMiddleware.RecoverHandler())

	docs.SwaggerInfo.Title = cfg.GetString(constants.ServiceName)
	docs.SwaggerInfo.Description = "Lists down the API to tackle ROBOT Apocalypse"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	healthCheckCtl := controllers.NewHealthCheckCtrl(l)
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/api/rob/v1/healthz", healthCheckCtl.HealthCheck)

	repo := repository.NewSurvivorRepo(l, dbIns)

	report := router.Group("/api/rob/v1/reports")
	{
		client := Client.NewExtClient(l, cfg)
		svc := services.NewReportSvc(repo, client)
		ctrl := controllers.NewReportCtrl(svc)

		report.GET("/list/robots", ctrl.ListAllRobots)
		report.GET("/list/survivors", ctrl.ListSurvivors)
		report.GET("/percentage/survivors", ctrl.ReportPercentage)
	}

	data := router.Group("/api/rob/v1/survivor")
	{
		svc := services.NewSurvivorSvc(repo, l)
		ctrl := controllers.NewController(svc)

		data.POST("/add", ctrl.AddSurvivor)
		data.PUT("/update", ctrl.UpdateSurvivorLocation)
		data.PUT("/flag", ctrl.FlagSurvivor)
	}

	return router
}
