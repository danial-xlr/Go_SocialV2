package ports

import (
	pb "common/proto/relation"
	"context"
	"relation/pkg/app"
	"relation/pkg/domain/relation"
)

type GrpcServer struct {
	pb.UnimplementedRelationServiceServer
	app app.Application
}

func NewGrpcServer(application app.Application) GrpcServer{
	return GrpcServer{app: application}
}

func (s GrpcServer)CreatRelation(ctx context.Context,r *pb.CreatRelationRequest)(*pb.CreatRelationResponse,error){
	rel:=relation.Relation{
		ProfileId: int(r.ProfileId),
		FollowingId: int(r.FollowingId),
	}
	err:=s.app.Commands.CreatRelation.Handle(ctx,&rel)
	if err != nil {
		return nil, err
	}
	return &pb.CreatRelationResponse{Relation: &pb.Relation{
		ProfileId: uint64(rel.ID),
		FollowingId: uint64(rel.FollowingId),
	}},nil
}

func (s GrpcServer)GetProfileFollowing(ctx context.Context,r *pb.GetProfileFollowingRequest)(*pb.GetProfileFollowingResponse,error){
	following:= []int{}
	err:=s.app.Queries.GetProfileFollowing.Handle(ctx,&following,int(r.ProfileId))
	if err != nil {
		return nil, err
	}
	f:= []uint64{}
	for _,a:=range following{
		f=append(f, uint64(a))
	}
	return &pb.GetProfileFollowingResponse{ProfileId: f},nil
}

func (s GrpcServer)GetProfileFollower(ctx context.Context,r *pb.GetProfileFollowerRequest)(*pb.GetProfileFollowerResponse,error){
	followers:=[]int{}
	err:=s.app.Queries.GetProfileFollowing.Handle(ctx,&followers,int(r.ProfileId))
	if err != nil {
		return nil, err
	}
	f:= []uint64{}
	for _,a:=range followers{
		f=append(f, uint64(a))
	}
	return &pb.GetProfileFollowerResponse{ProfileId: f},nil
}