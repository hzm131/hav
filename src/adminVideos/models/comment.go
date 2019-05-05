package models

import "github.com/jinzhu/gorm"

// 用户评论
type Comments struct {
	gorm.Model
	VideoId int `gorm:"column:video_id;type:integer;"json:"video_id"validate:"required||integer"` // 视频id 哪个视频的评论
	Content string `gorm:"column:content;type:varchar;"json:"content"validate:"required || string"` //评论内容
	PersonId int `gorm:"column:person_id;type:integer;"json:"person_id"validate:"required||integer"` // 评论人的id
	ToPersonId int `gorm:"column:to_person_id;type:integer;"json:"to_person_id"validate:"integer"` //评论目标用户id 没有为空
}



