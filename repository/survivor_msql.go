package repository

import (
	"database/sql"
	"github.com/RobotApocalypse/model"
	"github.com/gin-gonic/gin"
)

type MysqlSql struct {
	db *sql.DB
}

func (m MysqlSql) AddSurvivor(ctx *gin.Context, survivor *model.Survivor) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (m MysqlSql) AddSurvivorInventory(ctx *gin.Context, userId int64, inventory model.Inventory) error {
	//TODO implement me
	panic("implement me")
}

func (m MysqlSql) UpdateLocation(ctx *gin.Context, userId int64, location *model.Location) error {
	//TODO implement me
	panic("implement me")
}

func (m MysqlSql) FlagUser(ctx *gin.Context, userId int64) error {
	//TODO implement me
	panic("implement me")
}

func (m MysqlSql) FetchLastInsertedId(ctx *gin.Context) int64 {
	//TODO implement me
	panic("implement me")
}

func (m MysqlSql) CountInfectedLogs(ctx *gin.Context, userId int64) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (m MysqlSql) InsertInfectionLog(ctx *gin.Context, payload *model.FlagPayload) error {
	//TODO implement me
	panic("implement me")
}

func (m MysqlSql) CountSurvivors(ctx *gin.Context, infectionStatus int64) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (m MysqlSql) ListSurvivors(ctx *gin.Context, infectionStatus int64) ([]*model.Survivor, error) {
	//TODO implement me
	panic("implement me")
}

func NewSql(db *sql.DB) Survivor {
	return &MysqlSql{
		db: db,
	}
}
