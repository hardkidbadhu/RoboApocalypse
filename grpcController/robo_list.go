package grpcController

import (
	"context"

	robo "github.com/RobotApocalypse/grpc"
	"github.com/RobotApocalypse/services"
)

type server struct {
	robo.UnimplementedRoboApocalypseServer
	svc services.ReportService
}

func (s *server) GetRoboList(ctx context.Context, filter *robo.Filter) (*robo.RoboList, error) {
	robots, err := s.svc.ListRobots(ctx)
	if err != nil {
		return nil, err
	}

	roboIns := new(robo.RoboList)

	if filter.GetCategory() != "" {
		for i := range robots {
			if robots[i].Category == filter.GetCategory() {
				roboIns.RoboList = append(roboIns.RoboList, &robo.Robo{
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

		roboIns.RoboList = append(roboIns.RoboList, &robo.Robo{
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
