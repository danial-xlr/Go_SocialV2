package post

import "time"

type Post struct {
	ID        int    
	Body      string 
	ProfileId int    
	CreatAt   time.Time 
}
