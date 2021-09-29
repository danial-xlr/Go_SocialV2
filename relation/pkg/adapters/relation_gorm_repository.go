package adapters

import (
	"fmt"
	"relation/pkg/domain/relation"

	"github.com/jinzhu/gorm"
)

type MysqlRepository struct{
	DB *gorm.DB
}

type MySqlRelation struct{
	ID          int `json:"id"`
	ProfileId   int `json:"profile_id"`
	FollowingId int `json:"following_id"`
}

func NewGormRelationRepository(db *gorm.DB)MysqlRepository{
	if db==nil{
		panic("nil db")
	}
	db.AutoMigrate(&MySqlRelation{})
	return MysqlRepository{DB: db}
}

func (r MysqlRepository)CreatRelation(rel *relation.Relation) (error){
	err:=r.creatRelation(&MySqlRelation{
		ID: rel.ID,
		ProfileId: rel.ProfileId,
		FollowingId: rel.FollowingId,
	})
	if err!=nil{
		return err
	}
	return nil
}

func (r MysqlRepository)creatRelation(rel *MySqlRelation) error{
	err:=r.DB.Create(r).Error
	if err!=nil{
		return err
	}
	return nil
}

func (r MysqlRepository)GetProfileFollowing(f *[]int,profileId int) error{
	var rels []MySqlRelation
	err:=r.DB.Where("profile_id = ?",profileId).Find(&rels).Error
	if err!=nil{
		return err
	}
	for _,a := range rels {
		*f=append(*f,a.FollowingId)
	}
	return nil
}

func (r MysqlRepository)GetProfileFollower(f *[]int,profileId int) error{
	var rels []MySqlRelation
	err:=r.DB.Where("following_id = ?",profileId).Find(&rels).Error
	if err!=nil{
		return err
	}
	for _,a := range rels {
		*f=append(*f,a.ProfileId)
	}
	return nil
}

func DbURL()string{
	Host:=    "localhost"
  	Port:=     3306
	User:=     "root"
  	Password:= ""
  	DBName:=   "relation"
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		User,
		Password,
		Host,
		Port,
		DBName,
	)
	
}