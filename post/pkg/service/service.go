package service

import (
	"context"
	"os"
	"post/pkg/adapters"
	"post/pkg/app"
	"post/pkg/app/command"
	"post/pkg/app/query"

	//pb "common/proto/profile"
	cl "common/proto/profile"

	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
)

func NewApplication(ctx context.Context) app.Application {
	gormdb, err := gorm.Open("mysql", adapters.DbURL())
	if err != nil {
		panic("database connection error")
	}
	grpcAddr := os.Getenv("TRAINER_GRPC_ADDR")
	cli, _ := grpc.Dial(grpcAddr)
	c := cl.NewProfileServiceClient(cli)
	if err != nil {
		panic("client error")
	}
	profileClient := adapters.NewProfileGrpcClient(c)
	postRepo := adapters.NewMySqlPostRepository(gormdb)
	return app.Application{
		Commands: app.Commands{
			CreatPost: command.NewCreatPostHandler(postRepo, profileClient),
		},
		Queries: app.Queries{
			GetPostById:         query.NewGetPostByIdHandler(postRepo),
			GetPostsByProfileId: query.NewGetPostsByProfileIdHandler(postRepo),
		},
	}
}
