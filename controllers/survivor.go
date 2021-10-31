package controllers

import (
	"fmt"
	"github.com/RobotApocalypse/model"
	"github.com/RobotApocalypse/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Survivor interface {
	AddSurvivor(ctx *gin.Context)
	UpdateSurvivorLocation(ctx *gin.Context)
	FlagSurvivor(ctx *gin.Context)
}

type survivor struct {
	svc services.Survivor
}

// RobotApocalypse godoc
// @Tags Survivor
// @Summary add new survivor to the base
// @Accept json
// @Produce  json
// @Success 200 {object} model.Response
// @Failure 400 {string} json "{"message":0,"err_code": "ERR_*", "error": "err"}"
// @Failure 500 {string} json "{"message":0,"err_code": "ERR_*", "error": "err"}"
// @Param AddSurvivor body model.SurvivorPayload true "request body"
// @Router /api/rob/v1/survivor/add [POST]
func (s survivor) AddSurvivor(ctx *gin.Context) {
	survivorIns := model.SurvivorPayload{}
	if err := ctx.ShouldBindJSON(&survivorIns); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, model.ErrResp{
			Message: "invalid json",
			Err: err,
			ErrCode: model.ErrInvalidJSON,
		})
		return
	}

	if err := s.svc.AddSurvivor(ctx, &survivorIns); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err.(*model.ErrResp))
		return
	}

	ctx.IndentedJSON(http.StatusOK, model.Response{
		Message: fmt.Sprintf("Welcome aboard %s!!! ", survivorIns.Name),
	})
}

// RobotApocalypse godoc
// @Tags Survivor
// @Summary add new survivor to the base
// @Accept json
// @Produce  json
// @Success 200 {object} model.Response
// @Param userId query string true "user id"
// @Failure 400 {string} json "{"message":0,"err_code": "ERR_*", "error": "err"}"
// @Failure 500 {string} json "{"message":0,"err_code": "ERR_*", "error": "err"}"
// @Param UpdateSurvivorLocation body model.Location true "request body"
// @Router /api/rob/v1/survivor/update [PUT]
func (s survivor) UpdateSurvivorLocation(ctx *gin.Context) {
	survivorIns := model.Location{}
	if err := ctx.ShouldBindJSON(&survivorIns); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, model.ErrResp{
			Message: "invalid json",
			Err: err,
			ErrCode: model.ErrInvalidJSON,
		})
		return
	}

	if err := s.svc.UpdateSurvivorLocation(ctx, &survivorIns); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err.(*model.ErrResp))
		return
	}

	ctx.IndentedJSON(http.StatusOK, model.Response{
		Message: fmt.Sprintf("Location updated!!!"),
	})
}

// RobotApocalypse godoc
// @Tags Survivor
// @Summary add new survivor to the base
// @Accept json
// @Produce  json
// @Success 200 {object} model.Response
// @Failure 400 {string} json "{"message":0,"err_code": "ERR_*", "error": "err"}"
// @Failure 500 {string} json "{"message":0,"err_code": "ERR_*", "error": "err"}"
// @Param FlagSurvivor body model.FlagPayload true "request body"
// @Router /api/rob/v1/survivor/flag [PUT]
func (s survivor) FlagSurvivor(ctx *gin.Context) {
	survivorIns := model.FlagPayload{}
	if err := ctx.ShouldBindJSON(&survivorIns); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, model.ErrResp{
			Message: "invalid json",
			Err: err,
			ErrCode: model.ErrInvalidJSON,
		})
		return
	}

	markedInfected, err := s.svc.FlagSurvivor(ctx, &survivorIns)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err.(*model.ErrResp))
		return
	}

	if markedInfected {
		ctx.IndentedJSON(http.StatusOK, model.Response{
			Message: fmt.Sprintf("user already marked infected, thanks for your time!!!"),
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, model.Response{
		Message: fmt.Sprintf("user marked as infected!!!"),
	})
}

func NewController(svc services.Survivor) Survivor {
	return &survivor{
		svc: svc,
	}
}
