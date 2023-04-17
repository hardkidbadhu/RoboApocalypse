package repository

import (
	"github.com/hardkidbadhu/RoboApocalypse/model"

	"github.com/gin-gonic/gin"
)

type Survivor interface {
	AddSurvivor(ctx *gin.Context, survivor *model.Survivor) (int64, error)
	AddSurvivorInventory(ctx *gin.Context, userId int64, inventory model.Inventory) error
	UpdateLocation(ctx *gin.Context, userId int64, location *model.Location) error
	FlagUser(ctx *gin.Context, userId int64) error
	FetchLastInsertedId(ctx *gin.Context) int64
	CountInfectedLogs(ctx *gin.Context, userId int64) (int64, error)
	InsertInfectionLog(ctx *gin.Context, payload *model.FlagPayload) error
	CountSurvivors(ctx *gin.Context, infectionStatus int64) (int64, error)
	ListSurvivors(ctx *gin.Context, infectionStatus int64) ([]*model.Survivor, error)
}
