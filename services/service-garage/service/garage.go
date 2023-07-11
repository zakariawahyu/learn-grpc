package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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

func (g *GaragesServer) Add(ctx context.Context, param *model.GarageAndUserId) (*model.Garage, error) {
	userId := param.UserId
	garage := param.Garage

	if _, ok := localStorage.List[userId]; !ok {
		localStorage.List[userId] = new(model.GarageList)
		localStorage.List[userId].List = make([]*model.Garage, 0)
	}

	localStorage.List[userId].List = append(localStorage.List[userId].List, garage)

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.DataLoss, "Error metadata")
	}

	if v, ok := md["timestamp"]; ok {
		log.Println("Meta Data Timestamp")
		for i, val := range v {
			log.Println("%d - %s", i, val)
		}
	}

	log.Println("Adding garage", garage.String(), "for user", userId)

	return localStorage.List[userId].List[len(localStorage.List[userId].List)-1], nil
}

func (g *GaragesServer) List(ctx context.Context, param *model.GarageUserId) (*model.GarageList, error) {
	userId := param.UserId

	return localStorage.List[userId], nil
}
