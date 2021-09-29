package service

import (
	postpb "common/proto/post"
	realtionpb "common/proto/relation"
	"context"
	"feed/pkg/adapters"
	"feed/pkg/app"
	"feed/pkg/app/command"
	"feed/pkg/app/query"
	"os"

	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
)

func NewApplication(ctx context.Context)app.Application{
	gormdb,err:=gorm.Open("mysql",adapters.DbURL())
	if err!=nil{
		panic("database connection error")
	}
	pgrpcAddr := os.Getenv("TRAINER_GRPC_ADDR")
	pcli, _ := grpc.Dial(pgrpcAddr)
	pc := postpb.NewPostServiceClient(pcli)
	if err != nil {
		panic("client error")
	}
	postClient:=adapters.NewPostGrpcClient(pc)
	rgrpcAddr := os.Getenv("TRAINER_GRPC_ADDR")
	rcli, _ := grpc.Dial(rgrpcAddr)
	rc:=realtionpb.NewRelationServiceClient(rcli)
	relationClient:=adapters.NewRelationGrpcClient(rc)
	feedRepo:=adapters.NewMySqlFeedRepository(gormdb)
	return app.Application{
		Commands: app.Commands{
			CreatFeed: command.NewCreatFeedHandler(feedRepo,postClient,relationClient),
		},
		Queries: app.Queries{
			GetProfileFeed: query.NewGetProfileFeedHandler(feedRepo),
		},
	}
}