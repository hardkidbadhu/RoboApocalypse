package services

import (
	"context"
	"errors"
	"sort"
	"strconv"
	"sync"

	"github.com/RobotApocalypse/Client"
	"github.com/RobotApocalypse/model"
	"github.com/RobotApocalypse/repository"
	"github.com/gin-gonic/gin"
)

type ReportService interface {
	Percentage(ctx *gin.Context, filter string) (*model.PercentageResp, error)
	List(ctx *gin.Context, filter string) ([]*model.Survivor, error)
	ListRobots(ctx context.Context) ([]model.Robots, error)
}

type reportSvc struct {
	repo      repository.Survivor
	extClient Client.ExtClient
}

func (r reportSvc) Percentage(ctx *gin.Context, filter string) (*model.PercentageResp, error) {

	wt := sync.WaitGroup{}
	wt.Add(2)

	var (
		infected, total int64
		percentage      float64
		err             error
	)

	switch filter {
	case "infected":
		go func() {
			defer wt.Done()
			infected, err = r.repo.CountSurvivors(ctx, 2)
		}()

		go func() {
			defer wt.Done()
			total, err = r.repo.CountSurvivors(ctx, -1)
		}()
	case "nonInfected":
		go func() {
			defer wt.Done()
			infected, err = r.repo.CountSurvivors(ctx, 1)
		}()

		go func() {
			defer wt.Done()
			total, err = r.repo.CountSurvivors(ctx, -1)
		}()
	default:
		return nil, model.ErrResp{
			Err:     err,
			ErrCode: model.ErrInvalidFilter,
			Message: "invalid filter",
		}
	}

	wt.Wait()

	percentage = (float64(infected) / float64(total)) * 100
	return &model.PercentageResp{
		Percentage:      strconv.FormatFloat(percentage, 'f', 0, 64) + "%",
		InfectionStatus: filter,
	}, nil
}

func (r reportSvc) List(ctx *gin.Context, filter string) ([]*model.Survivor, error) {
	switch filter {
	case "infected":
		return r.repo.ListSurvivors(ctx, 2)
	case "nonInfected":
		return r.repo.ListSurvivors(ctx, 1)

	default:
		return nil, model.ErrResp{
			Err:     errors.New("invalid filter"),
			ErrCode: model.ErrInvalidFilter,
			Message: "invalid filter",
		}
	}
}

func (r reportSvc) ListRobots(ctx context.Context) ([]model.Robots, error) {
	robolist, err := r.extClient.FetchAllRobot(ctx)
	if err != nil {
		return nil, model.ErrResp{
			Err:     err,
			ErrCode: model.ErrInternalSRVError,
			Message: "error fetching data",
		}
	}
	sort.Sort(model.RoboList(robolist))
	return robolist, nil
}

func NewReportSvc(repo repository.Survivor, extClient Client.ExtClient) ReportService {
	return &reportSvc{
		repo:      repo,
		extClient: extClient,
	}
}
