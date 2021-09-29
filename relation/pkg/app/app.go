package app

import (
	"relation/pkg/app/command"
	"relation/pkg/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreatRelation command.CreatRelationHandler
}

type Queries struct {
	GetProfileFollower query.GetProfileFollowerHandler
	GetProfileFollowing query.GetProfileFollowingHandler
}