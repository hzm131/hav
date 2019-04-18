package main

import "github.com/jinzhu/gorm"

//评论
type Comment struct {
	gorm.Model
	PersonsId int  `gorm:"column:persons_id;type:integer;not null;"json:"persons_id"`  //这是谁的评论
	Content string `gorm:"column:content;type:integer;not null;"json:"content"` //评论内容
}

