package query

import (
	"context"
	"post/pkg/domain/post"
)

type GetPostsByProfileId struct{
	postRepo post.Repository
}

func NewGetPostsByProfileIdHandler(postRepo post.Repository)GetPostsByProfileId{
	if postRepo==nil{
		panic("nil profile repo")
	}
	return GetPostsByProfileId{postRepo: postRepo}
}

func (p GetPostsByProfileId)Handle(ctx context.Context,id int)(*[]post.Post,error){
	posts:=[]post.Post{}
	err:=p.postRepo.GetPostsByProfileId(ctx,&posts,id)
	if err != nil {
		return nil, err
	}
	return &posts,nil
}

