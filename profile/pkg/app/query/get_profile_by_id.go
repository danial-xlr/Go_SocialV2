package query

import (
	"context"
	"profile/pkg/domain/profile"
)

type GetProfileByid struct{
	profileRepo profile.Repository
}

func NewGetProfileByidHandler(profileRepo profile.Repository)GetProfileByid{
	if profileRepo==nil{
		panic("nil profile repo")
	}
	return GetProfileByid{profileRepo: profileRepo}
}

func (p GetProfileByid)Handle(ctx context.Context, id int) (*profile.Profile,error){
	profile,err:=p.profileRepo.GetProfileByid(ctx,id)
	if err != nil {
		return nil, err
	}
	return profile,nil
}