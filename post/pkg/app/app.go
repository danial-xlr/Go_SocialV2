package app

import (
	"post/pkg/app/command"
	"post/pkg/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreatPost command.CreatPostHandler
}

type Queries struct {
	GetPostById query.GetPostById
	GetPostsByProfileId query.GetPostsByProfileId
}