package service

import (
	"context"
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

func (u *UsersServer) Register(ctx context.Context, param *model.User) (*model.Empty, error) {
	localStorage.List = append(localStorage.List, param)

	log.Println("Registering User : ", param.String())

	return new(model.Empty), nil
}

func (u *UsersServer) List(ctx context.Context, void *model.Empty) (*model.UserList, error) {
	return localStorage, nil
}