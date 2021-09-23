package app

import (
	"profile/pkg/app/command"
	"profile/pkg/app/query"
)


type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreatProfile command.CreatProfileHandler
}

type Queries struct {
	GetProfileByid query.GetProfileByid
}