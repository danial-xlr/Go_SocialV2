package main

import (
	"context"
	"relation/pkg/ports"
	"relation/pkg/service"
	"common/server"
	"common/proto/relation"

	"google.golang.org/grpc"
)

func main(){
	ctx := context.Background()

	application := service.NewApplication(ctx)
	server.RunGRPCServer(func(server *grpc.Server) {
		svc := ports.NewGrpcServer(application)
		relation.RegisterRelationServiceServer(server, svc)
	})
}