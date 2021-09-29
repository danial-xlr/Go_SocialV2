package command

import (
	"context"
	
	"post/pkg/domain/post"
)

type CreatPostHandler struct {
	postRepo       post.Repository
	profileService ProfileService
}

func NewCreatPostHandler(postRepo post.Repository,proservice ProfileService) CreatPostHandler {
	if proservice==nil{
		panic("nil profile service")
	}
	if postRepo == nil {
		panic("nil post repo")
	}
	return CreatPostHandler{postRepo: postRepo,profileService: proservice}
}

func (p CreatPostHandler) Handle(ctx context.Context, po *post.Post) error {
	err:=p.profileService.ValidateProfileID(ctx,po.ProfileId)
	if err!=nil{
		panic("profile validate error")
	}
	return p.postRepo.CreatPost(ctx, po)
}
