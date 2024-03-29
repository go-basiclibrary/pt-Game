package datamodels

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Id             int            `gorm:"column:id;not null;primaryKey;autoIncrement;"`
	UserName       string         `gorm:"column:user_name;index"`
	Password       string         `gorm:"column:password"`
	Kind           int            `gorm:"column:kind"`
	Email          string         `gorm:"column:email;unique"`
	Gender         int            `gorm:"column:gender"`
	Country        string         `gorm:"column:country;index"`
	Region         string         `gorm:"column:region;index"`
	YoutubeAccount string         `gorm:"column:youtube_account"`
	YoutubeFansNum int            `gorm:"column:youtube_fans_num"`
	Language       string         `gorm:"column:language;default:en"`
	Level          string         `gorm:"column:level"`
	IsAdult        int            `gorm:"column:is_adult"`
	IsLocked       int            `gorm:"column:is_locked"`
	CreatedAt      time.Time      `gorm:"column:created_at;not null;index"`
	UpdatedAt      time.Time      `gorm:"column:updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at"`
	Updater        int            `gorm:"column:updater"`
}

func (User) TableName() string {
	return "streamer_user"
}
