package command

import (
	"context"
	"profile/pkg/domain/profile"
)


type CreatProfileHandler struct{
	profileRepo profile.Repository
}

func NewCreatProfileHandler(profileRepo profile.Repository) CreatProfileHandler {
	if profileRepo==nil{
		panic("nil profile repo")
	}
	return CreatProfileHandler{profileRepo: profileRepo}
}

func (p CreatProfileHandler) Handle(ctx context.Context, pro *profile.Profile) error{
	return p.profileRepo.CreatProfile(ctx,pro)
}
