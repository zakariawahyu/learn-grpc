package service

import (
	"context"
	"learn-grpc/model"
	"log"
)

var localStorage *model.GarageListBuyer

func init() {
	localStorage = new(model.GarageListBuyer)
	localStorage.List = make(map[string]*model.GarageList)
}

type GaragesServer struct {
	model.UnimplementedGaragesServer
}

func (g *GaragesServer) Add(ctx context.Context, param *model.GarageAndUserId) (*model.Empty, error) {
	userId := param.UserId
	garage := param.Garage

	if _, ok := localStorage.List[userId]; !ok {
		localStorage.List[userId] = new(model.GarageList)
		localStorage.List[userId].List = make([]*model.Garage, 0)
	}

	localStorage.List[userId].List = append(localStorage.List[userId].List, garage)

	log.Println("Adding garage", garage.String(), "for user", userId)

	return new(model.Empty), nil
}

func (g *GaragesServer) List(ctx context.Context, param *model.GarageUserId) (*model.GarageList, error) {
	userId := param.UserId

	return localStorage.List[userId], nil
}
