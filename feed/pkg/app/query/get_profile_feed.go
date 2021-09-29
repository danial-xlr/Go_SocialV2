package query

import (
	"context"
	"feed/pkg/domain/feed"
)

type GetProfileFeedHandler struct{
	feedRepo feed.Repository
}

func NewGetProfileFeedHandler(feedRepo feed.Repository)GetProfileFeedHandler{
	if feedRepo==nil{
		panic("nil feed repo")
	}
	return GetProfileFeedHandler{feedRepo: feedRepo}
}

func (f GetProfileFeedHandler)Handle(ctx context.Context,profileId uint64)(*feed.Feed,error){
	return f.feedRepo.GetProfileFeed(profileId)
}