package repository

import (
	"github.com/hardkidbadhu/RoboApocalypse/model"

	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// query
const (
	AddSurvivor         = `INSERT INTO tbl_survivor (name,age,gender,lastLocation) VALUES ($1,$2,$3,POINT($4,$5))`
	InsertInfectionLogs = `INSERT INTO tbl_infected_logs ("userId","infected","reportedBy") VALUES ($1,$2,$3)`

	AddSurvivorInventory = `INSERT INTO tbl_survivor_inventory (water,food,medication,ammunation,"userId") VALUES ($1,$2,$3,$4,$5)`

	UpdateLocation = `UPDATE tbl_survivor SET  lastLocation = POINT($1,$2) where id = $3`

	FlagUser = `UPDATE tbl_survivor SET  "infectionStatus" = $1 where id = $2`

	lastInsertedIdQuery = "select max(id) from tbl_survivor"

	CountInfectionLogs = `select count(*) from tbl_infected_logs where "userId"=$1 AND "infected"=1`

	CountInfectedSurvivors = `select count(*) from tbl_survivor where "infectionStatus"=$1`

	CountTotalSurvivors = `select count(*) from tbl_survivor`

	ListSurvivors = `select tbl_survivor.id, tbl_survivor.name, tbl_survivor.age, tbl_survivor.gender, tbl_survivor.lastlocation from tbl_survivor  where "infectionStatus"=$1`
)

// type Survivor interface {
// 	AddSurvivor(ctx *gin.Context, survivor *model.Survivor) (int64, error)
// 	AddSurvivorInventory(ctx *gin.Context, userId int64, inventory model.Inventory) error
// 	UpdateLocation(ctx *gin.Context, userId int64, location *model.Location) error
// 	FlagUser(ctx *gin.Context, userId int64) error
// 	FetchLastInsertedId(ctx *gin.Context) int64
// 	CountInfectedLogs(ctx *gin.Context, userId int64) (int64, error)
// 	InsertInfectionLog (ctx *gin.Context,  payload *model.FlagPayload) error
// 	CountSurvivors(ctx *gin.Context, infectionStatus int64) (int64, error)
// 	ListSurvivors(ctx *gin.Context, infectionStatus int64) ([]*model.Survivor, error)
// }

type survivorSvc struct {
	log *logrus.Entry
	db  *sql.DB
}

func (s survivorSvc) ListSurvivors(ctx *gin.Context, infectionStatus int64) ([]*model.Survivor, error) {
	var (
		rows         *sql.Rows
		err          error
		survivorList []*model.Survivor
	)

	switch infectionStatus {
	case 1:
		rows, err = s.db.QueryContext(ctx, ListSurvivors, 1)
	case 0:
		rows, err = s.db.QueryContext(ctx, ListSurvivors, 0)
	default:
		rows, err = s.db.QueryContext(ctx, CountTotalSurvivors)
	}

	if err != nil {
		s.log.Errorf("repo: ListSurvivors - %s", err.Error())
		return nil, err
	}

	if rows.Err() != nil {
		s.log.Errorf("repo: ListSurvivors (rows.Err()) - %s", err.Error())
		return nil, err
	}

	for rows.Next() {
		survivIns := model.Survivor{
			SurvivorPayload: new(model.SurvivorPayload),
		}
		if err := rows.Scan(
			&survivIns.Id,
			&survivIns.Name,
			&survivIns.Age,
			&survivIns.Gender,
			&survivIns.CoOrdinates,
		); err != nil {
			return nil, model.ErrResp{
				Err:     err,
				ErrCode: model.ErrDB,
				Message: "error in fetching data",
			}
		}

		survivorList = append(survivorList, &survivIns)
	}
	return survivorList, nil
}

func (s survivorSvc) CountSurvivors(ctx *gin.Context, infectionStatus int64) (int64, error) {
	var (
		count int64
		row   *sql.Row
	)

	switch infectionStatus {
	case 1:
		row = s.db.QueryRowContext(ctx, CountInfectedSurvivors, 1)
	case 0:
		row = s.db.QueryRowContext(ctx, CountInfectedSurvivors, 0)
	default:
		row = s.db.QueryRowContext(ctx, CountTotalSurvivors)
	}

	err := row.Scan(&count)
	if err != nil {
		s.log.Errorf("repo: CountSurvivors - %s", err.Error())
		return 0, err
	}

	return count, nil
}

func (s survivorSvc) CountInfectedLogs(ctx *gin.Context, userId int64) (int64, error) {
	var infectedLogs int64
	err := s.db.QueryRowContext(ctx, CountInfectionLogs, userId).Scan(&infectedLogs)
	if err != nil {
		s.log.Errorf("repo: CountInfectedLogs - %s", err.Error())
		return 0, err
	}

	return infectedLogs, nil
}

func (s survivorSvc) FetchLastInsertedId(ctx *gin.Context) int64 {
	var lastInsertedId int64
	err := s.db.QueryRowContext(ctx, lastInsertedIdQuery).Scan(&lastInsertedId)
	if err != nil {
		s.log.Errorf("repo: FetchLastInsertedId - %s", err.Error())
	}

	return lastInsertedId
}

func (s survivorSvc) AddSurvivor(ctx *gin.Context, survivor *model.Survivor) (int64, error) {
	_, err := s.db.ExecContext(
		ctx,
		AddSurvivor,
		survivor.Name,
		survivor.Age,
		survivor.Gender,
		survivor.Location.Latitude,
		survivor.Location.Longitude,
	)
	if err != nil {
		s.log.Errorf("repo: AddSurvivor - %s", err.Error())
		return 0, err
	}

	log.Println("id", s.FetchLastInsertedId(ctx))
	return s.FetchLastInsertedId(ctx), nil
}

func (s survivorSvc) AddSurvivorInventory(ctx *gin.Context, userId int64, inventory model.Inventory) error {
	s.log.Infoln("user id - %d", userId)
	s.log.Infoln("user id - %d", AddSurvivorInventory, userId)

	_, err := s.db.ExecContext(
		ctx,
		AddSurvivorInventory,
		inventory.Water,
		inventory.Food,
		inventory.Medication,
		inventory.Ammunition,
		userId,
	)
	if err != nil {
		s.log.Errorf("repo: AddSurvivorInventory - %s", err.Error())
		return err
	}

	return nil
}

func (s survivorSvc) UpdateLocation(ctx *gin.Context, userId int64, location *model.Location) error {
	_, err := s.db.ExecContext(
		ctx,
		UpdateLocation,
		location.Latitude,
		location.Longitude,
		userId,
	)

	if err != nil {
		s.log.Errorf("repo: Update - %s", err.Error())
		return err
	}

	return nil
}

func (s survivorSvc) InsertInfectionLog(ctx *gin.Context, payload *model.FlagPayload) error {
	_, err := s.db.ExecContext(
		ctx,
		InsertInfectionLogs,
		payload.InfectedUser,
		payload.InfectionStatus,
		payload.User,
	)

	if err != nil {
		s.log.Errorf("repo: InsertInfectionLog - %s", err.Error())
		return err
	}

	return nil
}

func (s survivorSvc) FlagUser(ctx *gin.Context, userId int64) error {
	_, err := s.db.ExecContext(
		ctx,
		FlagUser,
		2,
		userId,
	)

	if err != nil {
		s.log.Errorf("repo: FlagUser - %s", err.Error())
		return err
	}

	return nil
}

func NewSurvivorRepo(log *logrus.Entry, db *sql.DB) Survivor {
	return &survivorSvc{
		log: log,
		db:  db,
	}
}
