package ports

import (
	pb "common/proto/post"
	"context"
	"post/pkg/app"
	"post/pkg/domain/post"
	"time"
)

type GrpsServer struct {
	pb.UnimplementedPostServiceServer
	app app.Application
}

func NewGrpcServer(application app.Application) GrpsServer {
	return GrpsServer{app: application}
}

func (s GrpsServer) CreatPost(ctx context.Context, r *pb.CreatPostRequest) (*pb.CreatPostResponse, error) {
	post := post.Post{
		//ID: r.id,
		Body:      r.Body,
		ProfileId: int(r.ProfileId),
		CreatAt:   time.Now(),
	}
	err := s.app.Commands.CreatPost.Handle(ctx, &post)
	if err != nil {
		return nil, err
	}
	return &pb.CreatPostResponse{Post: &pb.Post{
		Id:        uint64(post.ID),
		Body:      post.Body,
		ProfileId: uint64(post.ProfileId),
		//CreatAt: post.CreatAt,
	}}, nil
}

func (s GrpsServer) GetPostById(ctx context.Context, r *pb.GetPostByIdRequest) (*pb.GetPostByIdResponse, error) {
	post, err := s.app.Queries.GetPostById.Handle(ctx, int(r.Id))
	if err != nil {
		return nil, err
	}
	return &pb.GetPostByIdResponse{Post: &pb.Post{
		Id:        uint64(post.ID),
		Body:      post.Body,
		ProfileId: uint64(post.ProfileId),
		//CreatAt: post.CreatAt,
	}}, nil
}

func (s GrpsServer) GetPostsByProfileId(ctx context.Context, r *pb.GetPostsByProfileIdRequest) (*pb.GetPostsByProfileIdResponse, error) {
	p, err := s.app.Queries.GetPostsByProfileId.Handle(ctx, int(r.ProfileId))
	if err != nil {
		return nil, err
	}
	posts := []*pb.Post{}
	for _, a := range *p {
		posts = append(posts, &pb.Post{
			Id:        uint64(a.ID),
			ProfileId: uint64(a.ProfileId),
			Body:      a.Body,
		})
	}
	return &pb.GetPostsByProfileIdResponse{Post: posts}, nil
}
