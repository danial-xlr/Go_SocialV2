package adapters

import (
	"feed/pkg/domain/feed"
	"fmt"

	"github.com/jinzhu/gorm"
)

type MySqlfeedRepository struct{
	DB *gorm.DB
}

type MySqlFeed struct{
	ID        	uint64 	 `json:"id"`
	ProfileId 	uint64	 `json:"profile_id"`
	Posts 		[]uint64 `json:"feed_post"`
}

func NewMySqlFeedRepository(db *gorm.DB)MySqlfeedRepository{
	if db==nil{
		panic("nil db")
	}
	db.AutoMigrate(&MySqlFeed{})
	return MySqlfeedRepository{DB: db}
}

func (m MySqlfeedRepository)CreatFeed(profileId uint64,posts []uint64)(*feed.Feed,error){
	f:=MySqlFeed{
		ProfileId: profileId,
		Posts: posts,
	}
	err:=m.DB.Create(f).Error
	if err!=nil{
		return nil,err
	}
	return &feed.Feed{ID: f.ID,ProfileId: f.ProfileId,Posts: f.Posts},nil
}

func (m MySqlfeedRepository)GetProfileFeed(profileId uint64)(*feed.Feed,error){
	f:=MySqlFeed{}
	err:=m.DB.Where("profile_id = ?",profileId).Find(&f).Error
	if err!=nil{
		return nil,err
	}
	return &feed.Feed{ID: f.ID,ProfileId: f.ProfileId,Posts: f.Posts},nil
}

//MySQl Config
func DbURL()string{
	Host:=    "localhost"
  	Port:=     3306
	User:=     "root"
  	Password:= ""
  	DBName:=   "feed"
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		User,
		Password,
		Host,
		Port,
		DBName,
	)
	
}