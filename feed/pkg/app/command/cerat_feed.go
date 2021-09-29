package command

import (
	"context"
	"feed/pkg/domain/feed"
)

type CreatFeedHandler struct{
	feedRepo feed.Repository
	postService PostService
	relationService RelationService
}

func NewCreatFeedHandler(feedRepo feed.Repository,postService PostService,relationService RelationService)CreatFeedHandler{
	if feedRepo==nil{
		panic("nil feed repo")
	}
	return CreatFeedHandler{feedRepo: feedRepo,postService: postService,relationService: relationService}
}

func (p CreatFeedHandler)Handle(ctx context.Context,profileId uint64)(*feed.Feed,error){
	followings,err:=p.relationService.GetProfileFollowing(ctx,profileId)
	if err!=nil{
		return nil,err	
	}
	followingPostsIds,err:=p.postService.GetFeedPosts(ctx,followings)
	if err!=nil{
		return nil,err	
	}

	return p.feedRepo.CreatFeed(profileId,followingPostsIds)
}