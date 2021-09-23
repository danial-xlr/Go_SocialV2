package query

import (
	"context"
	"post/pkg/domain/post"
)


type GetPostById struct{
	postRepo post.Repository
}

func NewGetPostByIdHandler(postRepo post.Repository)GetPostById{
	if postRepo==nil{
		panic("nil profile repo")
	}
	return GetPostById{postRepo: postRepo}
}

func (p GetPostById)Handle(ctx context.Context,id int)(*post.Post,error){
	post:=post.Post{}
	err:=p.postRepo.GetPostById(ctx,&post,id)
	if err != nil {
		return nil, err
	}
	return &post,nil
}