package query

import (
	"context"
	"relation/pkg/domain/relation"
)

type GetProfileFollowingHandler struct {
	relationRepo relation.Repository
}

func NewGetProfileFollowingHandler(relationRepo relation.Repository)GetProfileFollowingHandler{
	if relationRepo==nil{
		panic("nil profile repo")
	}
	return GetProfileFollowingHandler{relationRepo: relationRepo}
}

func (r GetProfileFollowingHandler)Handle(ctx context.Context,following *[]int,profileId int)(error){
	return r.relationRepo.GetProfileFollowing(following,profileId)
}