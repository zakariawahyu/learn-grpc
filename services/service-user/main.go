package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"learn-grpc/common/config"
	"learn-grpc/model"
	"learn-grpc/services/service-user/service"
	"log"
	"net"
)

func unaryServerInterceptorImpl(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Println("Incoming request", info.FullMethod)

	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "Error Invalid Argument")
	}

	if !isValidKey(meta["app_key"]) {
		return nil, status.Errorf(codes.Unauthenticated, "Unauthenticated")
	}

	md, err := handler(ctx, req)
	if err != nil {
		return nil, status.Error(codes.DataLoss, err.Error())
	}

	return md, nil
}

func isValidKey(appKey []string) bool {
	if len(appKey) < 1 {
		return false
	}

	return appKey[0] == "0723"
}

func main() {
	srv := grpc.NewServer(grpc.UnaryInterceptor(unaryServerInterceptorImpl))
	userSrv := service.UsersServer{}
	model.RegisterUsersServer(srv, &userSrv)

	log.Println("Starting RPC server at", config.SERVICE_USER_PORT)

	l, err := net.Listen("tcp", config.SERVICE_USER_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.SERVICE_USER_PORT, err)
	}

	log.Fatal(srv.Serve(l))
}
