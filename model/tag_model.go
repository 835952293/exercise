package model

import "time"

type Tag struct {
	Name       string    `json:"name"`
	State      uint8     `json:"stase"`
	CreateTime time.Time `gorm:"column:create_time;default:null" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;default:null" json:"update_time"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}
