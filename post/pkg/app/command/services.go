package command

import "context"


type ProfileService interface{
	ValidateProfileID(ctx context.Context,profileId int)error
}