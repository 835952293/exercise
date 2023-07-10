package model

import (
	"time"

	"gorm.io/datatypes"
)

type ArticleTag struct {
	TagId      uint32         `json:"tag_id"`
	ArticleId  uint32         `json:"article_id"`
	CreateTime time.Time      `gorm:"column:create_time;default:null" json:"create_time"`
	UpdateTime time.Time      `gorm:"column:update_time;default:null" json:"update_time"`
	JsonExtend datatypes.JSON `gorm:"column:json_extend" json:"json_extend"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}
