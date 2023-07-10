package main

import (
	"google.golang.org/grpc"
	"learn-grpc/common/config"
	"learn-grpc/model"
	"learn-grpc/services/service-garage/service"
	"log"
	"net"
)

func main() {
	srv := grpc.NewServer()
	garageSrv := service.GaragesServer{}
	model.RegisterGaragesServer(srv, &garageSrv)

	log.Println("Starting RPC server at", config.SERVICE_GARAGE_PORT)

	l, err := net.Listen("tcp", config.SERVICE_GARAGE_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.SERVICE_GARAGE_PORT, err)
	}

	log.Fatal(srv.Serve(l))
}
