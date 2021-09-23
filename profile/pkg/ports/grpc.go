package ports

import (
	pb "common/proto/profile"
	"context"
	"profile/pkg/app"
	"profile/pkg/domain/profile"
)

type GrpcServer struct {
	app app.Application
}

func NewGrpcServer(application app.Application) GrpcServer {
	return GrpcServer{app: application}
}

func (s GrpcServer) PostProfile(ctx context.Context, r *pb.PostProfileRequest) (*pb.PostProfileResponse, error) {
	p := profile.Profile{
		UserName: r.UserName,
	}
	err := s.app.Commands.CreatProfile.Handle(ctx, &p)
	if err != nil {
		return nil, err
	}
	return &pb.PostProfileResponse{Profile: &pb.Profile{
		Id:       uint64(p.ID),
		UserName: p.UserName,
	}}, nil
}

func (s GrpcServer) GetProfile(ctx context.Context, r *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {
	p, err := s.app.Queries.GetProfileByid.Handle(ctx, int(r.Id))
	if err != nil {
		return nil, err
	}
	return &pb.GetProfileResponse{Profile: &pb.Profile{
		Id:       uint64(p.ID),
		UserName: p.UserName,
	}}, nil
}
