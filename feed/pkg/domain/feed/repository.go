package feed

type Repository interface{
	CreatFeed(profileId uint64,posts []uint64)(*Feed,error)
	GetProfileFeed(profileId uint64)(*Feed,error)
}

