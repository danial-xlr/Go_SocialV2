package profile

import "context"

type Repository interface {	
	CreatProfile(ctx context.Context, p *Profile) (err error)	
	GetProfileByid(ctx context.Context, id int) (p *Profile,err error)	
}