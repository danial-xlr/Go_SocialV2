package adapters

import (	
	pb "common/proto/profile"
	"context"
)
type profileGrpcClient struct {
	client pb.ProfileServiceClient
}


func NewProfileGrpcClient(client pb.ProfileServiceClient) profileGrpcClient{
	return profileGrpcClient{client: client}
}

func (p profileGrpcClient)ValidateProfileID(ctx context.Context,profileId int)error{
	_,err:=p.client.GetProfile(ctx,&pb.GetProfileRequest{Id: uint64(profileId)})
	if err!=nil{
		return err
	}
	return nil
}