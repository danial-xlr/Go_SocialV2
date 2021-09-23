package main

import (
	"context"
	"profile/pkg/ports"
	"profile/pkg/service"
	"common/server"
	"common/proto/profile"

	"google.golang.org/grpc"
)

func main() {	
	ctx := context.Background()

	application := service.NewApplication(ctx)

	
	
	server.RunGRPCServer(func(server *grpc.Server) {
		svc := ports.NewGrpcServer(application)
		profile.RegisterProfileServiceServer(server, svc)
	})
	
}	
