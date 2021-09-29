package relation


type Repository interface {	
	CreatRelation(r *Relation) (error)
	GetProfileFollowing(f *[]int,profileId int) (error)
	GetProfileFollower(f *[]int,profileId int) (error)
}