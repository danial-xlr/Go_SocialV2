package service

import (
	"context"
	"post/pkg/adapters"
	"post/pkg/app"
	"post/pkg/app/command"
	"post/pkg/app/query"

	"github.com/jinzhu/gorm"
)

func NewApplication(ctx context.Context)app.Application{
	gormdb,err:=gorm.Open("mysql",adapters.DbURL())
	if err!=nil{
		panic("database connection error")
	}
	postRepo:=adapters.NewMySqlPostRepository(gormdb)
	return app.Application{
		Commands: app.Commands{
			CreatPost: command.NewCreatPostHandler(postRepo),
		},
		Queries: app.Queries{
			GetPostById: query.NewGetPostByIdHandler(postRepo),
			GetPostsByProfileId: query.NewGetPostsByProfileIdHandler(postRepo),
		},
	}
}