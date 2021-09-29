package command

import "context"

type PostService interface {
	GetFeedPosts(ctx context.Context,followingIds []uint64)([]uint64,error)
}

type RelationService interface{
	GetProfileFollowing(ctx context.Context,ProfileId uint64)([]uint64,error)
}