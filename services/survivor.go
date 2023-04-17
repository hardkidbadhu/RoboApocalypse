package services

import (
	"github.com/hardkidbadhu/RoboApocalypse/model"
	"github.com/hardkidbadhu/RoboApocalypse/repository"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Survivor interface {
	AddSurvivor(ctx *gin.Context, sIns *model.SurvivorPayload) error
	UpdateSurvivorLocation(ctx *gin.Context, location *model.Location) error
	FlagSurvivor(ctx *gin.Context, flagIns *model.FlagPayload) (bool, error)
}

type survivorSvc struct {
	repo repository.Survivor
	log  *logrus.Entry
}

func (s survivorSvc) AddSurvivor(ctx *gin.Context, sIns *model.SurvivorPayload) error {
	//1. Add the survivor to the base
	sId, err := s.repo.AddSurvivor(ctx, &model.Survivor{
		SurvivorPayload: sIns,
	})

	if err != nil {
		s.log.Errorf("svc: AddSurvivor (adding survivor) - %s", err.Error())
		return &model.ErrResp{
			Err:     err,
			ErrCode: model.ErrDB,
			Message: "error in inserting survivor data",
		}
	}

	err = s.repo.AddSurvivorInventory(ctx, sId, sIns.Inventory)
	if err != nil {
		return &model.ErrResp{
			Err:     err,
			ErrCode: model.ErrDB,
			Message: "error in inserting survivor's inventory data",
		}
	}

	return nil
}

func (s survivorSvc) UpdateSurvivorLocation(ctx *gin.Context, location *model.Location) error {
	userIdStr := ctx.Request.URL.Query().Get("userId")

	userId, _ := strconv.Atoi(userIdStr)
	if userId == 0 {
		return &model.ErrResp{
			ErrCode: model.ErrInvalidUSERID,
			Message: "invalid user id",
		}
	}

	err := s.repo.UpdateLocation(ctx, int64(userId), location)
	if err != nil {
		return &model.ErrResp{
			Err:     err,
			ErrCode: model.ErrDB,
			Message: "error in inserting survivor's inventory data",
		}
	}

	return nil
}

func (s survivorSvc) FlagSurvivor(ctx *gin.Context, flagIns *model.FlagPayload) (bool, error) {

	count, err := s.repo.CountInfectedLogs(ctx, int64(flagIns.InfectedUser))
	if err != nil {
		return false, &model.ErrResp{
			Err:     err,
			ErrCode: model.ErrDB,
			Message: "error in marking user infected, please try after sometime",
		}
	}

	switch count {
	case 2:
		if err := s.repo.InsertInfectionLog(ctx, flagIns); err != nil {
			return false, &model.ErrResp{
				Err:     err,
				ErrCode: model.ErrDB,
				Message: "error in marking user infected, please try after sometime",
			}
		}

		if err := s.repo.FlagUser(ctx, int64(flagIns.InfectedUser)); err != nil {
			return false, &model.ErrResp{
				Err:     err,
				ErrCode: model.ErrDB,
				Message: "error in marking user infected, please try after sometime",
			}
		}
	case 3:
		return true, nil
	default:
		if err := s.repo.InsertInfectionLog(ctx, flagIns); err != nil {
			return false, &model.ErrResp{
				Err:     err,
				ErrCode: model.ErrDB,
				Message: "error in marking user infected, please try after sometime",
			}
		}
	}

	return false, nil
}

func NewSurvivorSvc(repo repository.Survivor, log *logrus.Entry) Survivor {
	return &survivorSvc{
		repo: repo,
		log:  log,
	}
}
