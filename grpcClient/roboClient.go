package grpcclient

import (
	"context"

	robo "github.com/RobotApocalypse/grpc"
	"google.golang.org/grpc"
)

type RoboClient interface {
	GetRoboList(ctx context.Context, filter *robo.Filter) (*robo.RoboList, error)
}

type roboClient struct {
	conn   *grpc.ClientConn
	client robo.RoboApocalypseClient
}

func NewRoboClient(address string) (RoboClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := robo.NewRoboApocalypseClient(conn)

	return &roboClient{conn: conn, client: client}, nil
}

func (c roboClient) GetRoboList(ctx context.Context, filter *robo.Filter) (*robo.RoboList, error) {
	return c.client.GetRoboList(ctx, &robo.Filter{
		Category: filter.Category,
	})
}
