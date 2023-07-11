package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"learn-grpc/model"
	"log"
)

var localStorage *model.UserList

func init() {
	localStorage = new(model.UserList)
	localStorage.List = make([]*model.User, 0)
}

type UsersServer struct {
	model.UnimplementedUsersServer
}

func (u *UsersServer) Register(ctx context.Context, param *model.User) (*model.User, error) {
	localStorage.List = append(localStorage.List, param)

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

	log.Println("Registering User : ", param.String())

	return localStorage.List[len(localStorage.List)-1], nil
}

func (u *UsersServer) List(ctx context.Context, void *model.Empty) (*model.UserList, error) {
	return localStorage, nil
}
