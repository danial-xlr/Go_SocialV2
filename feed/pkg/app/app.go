package app

import (
	"feed/pkg/app/command"
	"feed/pkg/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreatFeed command.CreatFeedHandler
}

type Queries struct {
	GetProfileFeed query.GetProfileFeedHandler
}