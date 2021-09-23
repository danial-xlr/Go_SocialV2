package command

import (
	"context"
	"post/pkg/domain/post"
)

type CreatPostHandler struct{
	postRepo post.Repository
}

func NewCreatPostHandler(postRepo post.Repository)CreatPostHandler{
	if postRepo==nil{
		panic("nil post repo")
	}
	return CreatPostHandler{postRepo: postRepo}
}

func (p CreatPostHandler)Handle(ctx context.Context,po *post.Post)error{
	return p.postRepo.CreatPost(ctx,po)
}