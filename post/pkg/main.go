package main

import (
	"common/server"
	"context"
	
	"post/pkg/ports"
	"post/pkg/service"
	"common/proto/post"

	"google.golang.org/grpc"
)


func main(){
	ctx:=context.Background()
	application:=service.NewApplication(ctx)
	server.RunGRPCServer(func(server *grpc.Server) {
		svc:=ports.NewGrpcServer(application)
		post.RegisterPostServiceServer(server, svc)
	})
}