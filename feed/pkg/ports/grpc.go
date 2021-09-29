package ports

import (
	pb "common/proto/feed"
	"context"
	"feed/pkg/app"
)

type GrpcServer struct {
	pb.UnimplementedFeedServiceServer
	app app.Application
}

func NewGrpcServer(application app.Application) GrpcServer {
	return GrpcServer{app: application}
}

func (s GrpcServer) CreatFeed(ctx context.Context, r *pb.CreatFeedRequest) (*pb.CreatFeedResponse, error) {
	f, err := s.app.Commands.CreatFeed.Handle(ctx, r.ProfileId)
	if err != nil {
		return nil, err
	}
	return &pb.CreatFeedResponse{Feed: &pb.Feed{Id: f.ID, ProfileId: f.ProfileId, FeedPost: f.Posts}}, nil
}

func (s GrpcServer) GetProfileFeed(ctx context.Context, r *pb.GetProfileFeedRequest) (*pb.GetProfileFeedResponse, error) {
	f, err := s.app.Queries.GetProfileFeed.Handle(ctx, r.ProfileId)
	if err != nil {
		return nil, err
	}
	return &pb.GetProfileFeedResponse{Feed: &pb.Feed{Id: f.ID, ProfileId: f.ProfileId, FeedPost: f.Posts}}, nil
}
