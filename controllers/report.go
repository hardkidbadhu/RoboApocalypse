package controllers

import (
	"github.com/RobotApocalypse/model"
	"github.com/RobotApocalypse/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

type ReportController interface {
	ListAllRobots(ctx *gin.Context)
	ListSurvivors(ctx *gin.Context)
	ReportPercentage(ctx *gin.Context)
}

type reportCtrl struct {
	svc services.ReportService
}

// RobotApocalypse godoc
// @Tags Reports
// @Summary lists all the robots
// @Produce  json
// @Success 200 {array} model.Survivor
// @Param infectionStatus query string true "infected / nonInfected"
// @Failure 500 {string} json "{"message":0,"err_code": "ERR_*", "error": "err"}"
// @Router /api/rob/v1/reports/percentage/survivors [GET]
func (r reportCtrl) ReportPercentage(ctx *gin.Context) {
	filter := ctx.Request.URL.Query().Get("infectionStatus")
	log.Println("filter" ,filter, ctx.Request.URL)
	resp , err := r.svc.Percentage(ctx, strings.TrimSpace(filter))
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err.(model.ErrResp))
		return
	}

	ctx.IndentedJSON(http.StatusOK, resp)
}

// RobotApocalypse godoc
// @Tags Reports
// @Summary lists all the survivors
// @Produce  json
// @Success 200 {array} model.Survivor
// @Param infectionStatus query string true "infected / nonInfected"
// @Failure 500 {string} json "{"message":0,"err_code": "ERR_*", "error": "err"}"
// @Router /api/rob/v1/reports/list/survivors [GET]
func (r reportCtrl) ListSurvivors(ctx *gin.Context) {
	filter := ctx.Request.URL.Query().Get("infectionStatus")
	log.Println("filter" ,filter, ctx.Request.URL)
	resp , err := r.svc.List(ctx, strings.TrimSpace(filter))
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err.(model.ErrResp))
		return
	}

	ctx.IndentedJSON(http.StatusOK, resp)
}

// RobotApocalypse godoc
// @Tags Reports
// @Summary lists all the robots
// @Produce  json
// @Success 200 {array} model.Robots
// @Failure 500 {string} json "{"message":0,"err_code": "ERR_*", "error": "err"}"
// @Router /api/rob/v1/reports/list/robots [GET]
func (r reportCtrl) ListAllRobots(ctx *gin.Context) {
	robots, err := r.svc.ListRobots(ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err.(*model.ErrResp))
		return
	}
	ctx.IndentedJSON(http.StatusOK, robots)
}

func NewReportCtrl(svc services.ReportService) ReportController {
	return &reportCtrl{
		svc:svc,
	}
}