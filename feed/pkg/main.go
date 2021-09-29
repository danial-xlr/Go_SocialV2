package main

import (
	"context"
	"common/proto/feed"
	"feed/pkg/ports"
	"feed/pkg/service"
	"common/server"

	"google.golang.org/grpc"
)

func main(){
	ctx := context.Background()

	application := service.NewApplication(ctx)
	server.RunGRPCServer(func(server *grpc.Server) {
		svc := ports.NewGrpcServer(application)		
		feed.RegisterFeedServiceServer(server,svc)		
	})
}

