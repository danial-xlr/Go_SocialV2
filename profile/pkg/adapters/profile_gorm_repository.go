package adapters

import (
	"context"
	"fmt"
	"profile/pkg/domain/profile"

	"github.com/jinzhu/gorm"
)

type GormProfileRepository struct {
	DB *gorm.DB
}

type MySqlProfile struct {
	ID       int    `json:"id"`
	UserName string `json:"user_name"`
}

func NewGormProfileRepository(db *gorm.DB) GormProfileRepository {
	if db == nil {
		panic("nil db")
	}
	db.AutoMigrate(&MySqlProfile{})
	return GormProfileRepository{DB: db}
}

func (d GormProfileRepository) CreatProfile(ctx context.Context, p *profile.Profile) (err error) {
	err = d.creatProfile(&MySqlProfile{
		ID:       p.ID,
		UserName: p.UserName,
	})
	if err != nil {
		return err
	}
	return nil
}

func (d GormProfileRepository) GetProfileByid(ctx context.Context, id int) (p *profile.Profile, err error) {
	pro := MySqlProfile{}
	err = d.getProfileByid(&pro, id)
	if err != nil {
		return nil, err
	}
	return &profile.Profile{ID: pro.ID, UserName: pro.UserName}, nil
}

func (d GormProfileRepository) creatProfile(p *MySqlProfile) (err error) {
	err = d.DB.Create(p).Error
	if err != nil {
		return err
	}
	return nil
}

func (d GormProfileRepository) getProfileByid(p *MySqlProfile, id int) (err error) {
	err = d.DB.Where("id = ?", id).First(p).Error
	if err != nil {
		return err
	}
	return nil
}

//MySQl Config
func DbURL() string {
	Host := "localhost"
	Port := 3306
	User := "danial"
	Password := "123456"
	DBName := "profile"
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		User,
		Password,
		Host,
		Port,
		DBName,
	)

}
