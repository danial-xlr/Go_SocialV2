package query

import (
	"context"
	"relation/pkg/domain/relation"
)

type GetProfileFollowerHandler struct {
	relationRepo relation.Repository
}

func NewGetProfileFollowerHandler(relationRepo relation.Repository)GetProfileFollowerHandler{
	if relationRepo==nil{
		panic("nil profile repo")
	}
	return GetProfileFollowerHandler{relationRepo: relationRepo}
}

func (r GetProfileFollowerHandler)Handle(ctx context.Context,followers *[]int,profileId int)error{
	return r.relationRepo.GetProfileFollower(followers,profileId)
}