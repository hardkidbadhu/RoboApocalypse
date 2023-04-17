package Client

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/hardkidbadhu/RoboApocalypse/configuration"
	"github.com/hardkidbadhu/RoboApocalypse/constants"
	"github.com/hardkidbadhu/RoboApocalypse/model"
	"github.com/sirupsen/logrus"
)

type ExtClient interface {
	FetchAllRobot(ctx context.Context) ([]model.Robots, error)
}

type extClient struct {
	log *logrus.Entry
	cfg configuration.Config
}

func (e extClient) FetchAllRobot(ctx context.Context) ([]model.Robots, error) {
	req, err := http.NewRequest(http.MethodGet, e.cfg.GetString(constants.RoboListURL), nil)
	if err != nil {
		e.log.Errorf("extClient: FetchAllRobot (NewRequest) - %s", err.Error())
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		e.log.Errorf("extClient: FetchAllRobot (Do) - %s", err.Error())
		return nil, err
	}

	var robotList []model.Robots
	byt, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(byt, &robotList); err != nil {
		e.log.Errorf("extClient: FetchAllRobot (Unmarshal) - %s", err.Error())
		return nil, err
	}

	return robotList, nil
}

func NewExtClient(log *logrus.Entry, cfg configuration.Config) ExtClient {
	return &extClient{
		log: log,
		cfg: cfg,
	}
}
