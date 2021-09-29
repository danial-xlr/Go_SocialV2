package service

import (
	"context"
	"os"
	"relation/pkg/adapters"
	"relation/pkg/app"
	"relation/pkg/app/command"
	"relation/pkg/app/query"
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
	relationRepo:=adapters.NewGormRelationRepository(gormdb)
	return app.Application{
		Commands: app.Commands{
			CreatRelation: command.NewCreatRelationHandler(relationRepo,profileClient),
		},
		Queries: app.Queries{
			GetProfileFollower: query.NewGetProfileFollowerHandler(relationRepo),
			GetProfileFollowing: query.NewGetProfileFollowingHandler(relationRepo),
		},
		
	}
}