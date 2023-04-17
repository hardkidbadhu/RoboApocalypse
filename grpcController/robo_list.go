package grpcController

import (
	"context"
	"github.com/hardkidbadhu/RoboApocalypse/roboGrpc"

	"github.com/hardkidbadhu/RoboApocalypse/services"
)

type server struct {
	roboGrpc.UnimplementedRoboApocalypseServer
	svc services.ReportService
}

func (s *server) GetRoboList(ctx context.Context, filter *roboGrpc.Filter) (*roboGrpc.RoboList, error) {
	robots, err := s.svc.ListRobots(ctx)
	if err != nil {
		return nil, err
	}

	roboIns := new(roboGrpc.RoboList)

	if filter.GetCategory() != "" {
		for i := range robots {
			if robots[i].Category == filter.GetCategory() {
				roboIns.RoboList = append(roboIns.RoboList, &roboGrpc.Robo{
					Model:            robots[i].Model,
					SerialNumber:     robots[i].SerialNumber,
					ManufacturedDate: robots[i].ManufacturedDate.String(),
					Category:         robots[i].Category,
				})
			}
		}
		return roboIns, nil
	}

	for i := range robots {

		roboIns.RoboList = append(roboIns.RoboList, &roboGrpc.Robo{
			Model:            robots[i].Model,
			SerialNumber:     robots[i].SerialNumber,
			ManufacturedDate: robots[i].ManufacturedDate.String(),
			Category:         robots[i].Category,
		})
	}

	return roboIns, nil
}

func NewRoboServer(svc services.ReportService) *server {
	return &server{
		svc: svc,
	}
}
