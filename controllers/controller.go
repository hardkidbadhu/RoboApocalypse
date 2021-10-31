package controllers

import (
	"github.com/RobotApocalypse/constants"
	"github.com/RobotApocalypse/model"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type HealthCheckController struct {
	log *logrus.Entry
}

func NewHealthCheckCtrl(log *logrus.Entry) *HealthCheckController {
	return &HealthCheckController{
		log: log,
	}
}

// RobotApocalypse godoc
// @Tags Health
// @Summary Health
// @Description Health Check API
// @Produce  json
// @Success 200 {object} model.HealthResponse
// @Router /api/rob/v1/healthz [get]
func (healthCheckController HealthCheckController) HealthCheck(ctx *gin.Context) {
	logger := healthCheckController.log.
		WithField("Class", "Controller").
		WithField("Method", "HealthCheck")
	logger.Info("Health check method Initiated")
	ctx.JSON(http.StatusOK, model.HealthResponse{
		Status:  "UP",
		Version: constants.AppVersion,
	})
	logger.Info(" HealthCheck call completed")
}


