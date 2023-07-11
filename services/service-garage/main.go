package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"learn-grpc/common/config"
	"learn-grpc/model"
	"learn-grpc/services/service-garage/service"
	"log"
	"net"
)

func unaryServerInterceptorImpl(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Println("Incoming request", info.FullMethod)

	md, err := handler(ctx, req)
	if err != nil {
		return nil, status.Error(codes.DataLoss, err.Error())
	}

	return md, nil
}

func main() {
	srv := grpc.NewServer(grpc.UnaryInterceptor(unaryServerInterceptorImpl))
	garageSrv := service.GaragesServer{}
	model.RegisterGaragesServer(srv, &garageSrv)

	log.Println("Starting RPC server at", config.SERVICE_GARAGE_PORT)

	l, err := net.Listen("tcp", config.SERVICE_GARAGE_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.SERVICE_GARAGE_PORT, err)
	}

	log.Fatal(srv.Serve(l))
}
