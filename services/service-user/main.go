package main

import (
	"google.golang.org/grpc"
	"learn-grpc/common/config"
	"learn-grpc/model"
	"learn-grpc/services/service-user/service"
	"log"
	"net"
)

func main() {
	srv := grpc.NewServer()
	userSrv := service.UsersServer{}
	model.RegisterUsersServer(srv, &userSrv)

	log.Println("Starting RPC server at", config.SERVICE_USER_PORT)

	l, err := net.Listen("tcp", config.SERVICE_USER_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.SERVICE_USER_PORT, err)
	}

	log.Fatal(srv.Serve(l))
}
