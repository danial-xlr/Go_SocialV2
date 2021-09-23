package service

import (
	"context"
	"profile/pkg/adapters"
	"profile/pkg/app"
	"profile/pkg/app/command"
	"profile/pkg/app/query"

	"github.com/jinzhu/gorm"
)

func NewApplication(ctx context.Context) app.Application{
	gormdb,err:=gorm.Open("mysql",adapters.DbURL())
	if err!=nil{
		panic("database connection error")
	}
	profileRepo:=adapters.NewGormProfileRepository(gormdb)
	return app.Application{
		Commands: app.Commands{
			CreatProfile: command.NewCreatProfileHandler(profileRepo),
		},
		Queries: app.Queries{
			GetProfileByid: query.NewGetProfileByidHandler(profileRepo),
		},
	}
}