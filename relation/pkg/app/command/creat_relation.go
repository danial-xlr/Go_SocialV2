package command

import (
	"context"
	"relation/pkg/domain/relation"
)

type CreatRelationHandler struct {
	relationRepo relation.Repository
	profileService ProfileService
}

func NewCreatRelationHandler (relationRepo relation.Repository,proservice ProfileService) CreatRelationHandler{
	if proservice==nil{
		panic("nil profile service")
	}
	if relationRepo==nil{
		panic("nil profile repo")
	}
	return CreatRelationHandler{relationRepo: relationRepo,profileService: proservice}
}

func (r CreatRelationHandler)Handle(ctx context.Context , rel *relation.Relation )error{
	err:=r.profileService.ValidateProfileID(ctx,rel.ProfileId)
	if err!=nil{
		panic("profile validate error")
	}
	err=r.profileService.ValidateProfileID(ctx,rel.FollowingId)
	if err!=nil{
		panic("following Id validate error")
	}
	return r.relationRepo.CreatRelation(rel)
}
