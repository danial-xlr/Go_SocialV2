package adapters

import (
	postpb "common/proto/post"
	realtionpb "common/proto/relation"
	"context"
)

type postGrpcClient struct{
	client postpb.PostServiceClient
}

type relationGrpcClient struct {
	client realtionpb.RelationServiceClient
}

func NewPostGrpcClient(client postpb.PostServiceClient)postGrpcClient{
	return postGrpcClient{client: client}
}

func (p postGrpcClient)GetFeedPosts(ctx context.Context,followingIds []uint64)([]uint64,error){
	posts:=[]uint64{}
	for _,f:=range followingIds{
		postRes,err:=p.client.GetPostsByProfileId(ctx,&postpb.GetPostsByProfileIdRequest{ProfileId: f})
		if err!=nil{
			return nil,err
		}
		for _,a:=range postRes.Post{
			posts = append(posts, a.Id)
		}
	}
	return posts,nil
}

func NewRelationGrpcClient(client realtionpb.RelationServiceClient)relationGrpcClient{
	return relationGrpcClient{client: client}
}

func (r relationGrpcClient)GetProfileFollowing(ctx context.Context,ProfileId uint64)([]uint64,error){
	return r.GetProfileFollowing(ctx,ProfileId)
}