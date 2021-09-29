package adapters

import (
	"context"
	"fmt"
	"post/pkg/domain/post"
	"time"

	"github.com/jinzhu/gorm"
)

type MySqlRepository struct{
	DB *gorm.DB
}
type MySqlPost struct {
	ID        int    	`json:"id"`
	Body      string 	`json:"body"`
	ProfileId int    	`json:"profile_id"`
	CreatAt   time.Time `json:"creat_at"`
}

func NewMySqlPostRepository(db *gorm.DB)MySqlRepository{
	if db==nil{
		panic("nil db")
	}
	db.AutoMigrate(&MySqlPost{})
	return MySqlRepository{DB: db}
}

func (m MySqlRepository )CreatPost(ctx context.Context,p *post.Post) error {
	post:=MySqlPost{
		ID: p.ID,
		Body: p.Body,
		ProfileId: p.ProfileId,
		CreatAt: p.CreatAt,
	}
	err:=m.creatPost(&post)
	if err!=nil{
		return err
	}
	return nil
}

func (m MySqlRepository )GetPostById(ctx context.Context,p *post.Post,id int)error{
	post:=MySqlPost{}
	err:=m.getPostById(&post,id)
	if err!=nil{
		return err
	}
	return nil
}

func (m MySqlRepository )GetPostsByProfileId(ctx context.Context,p *[]post.Post,profileId int)error{
	posts:=[]MySqlPost{}
	err:=m.getPostsByProfileId(&posts,profileId)
	if err!=nil{
		return err
	}
	return nil
}

func (m MySqlRepository )creatPost(p *MySqlPost) error {
	err:=m.DB.Create(p).Error
	if err!=nil{
		return err
	}
	return nil
}

func (m MySqlRepository )getPostById(p *MySqlPost,id int)error{
	err:=m.DB.Where("id = ?",id).First(p).Error
	if err!=nil{
		return err
	}
	return nil
}

func (m MySqlRepository )getPostsByProfileId(p *[]MySqlPost,profileId int)error{
	err:=m.DB.Where("profile_id = ?",profileId).Find(p).Error
	if err!=nil{
		return err
	}
	return nil
}

//MySQl Config
func DbURL()string{
	Host:=    "localhost"
  	Port:=     3306
	User:=     "root"
  	Password:= ""
  	DBName:=   "post"
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		User,
		Password,
		Host,
		Port,
		DBName,
	)
	
}