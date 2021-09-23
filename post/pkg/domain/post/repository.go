package post

import "context"

type Repository interface{
	CreatPost(ctx context.Context,p *Post)error
	GetPostById(ctx context.Context,p *Post,id int)error
	GetPostsByProfileId(ctx context.Context,p *[]Post,profileId int)error
}