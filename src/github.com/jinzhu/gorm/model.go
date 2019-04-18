package gorm

import "time"

// Model base model definition, including fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`, which could be embedded in your models
//    type User struct {
//      gorm.Model
//    }
type Model struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time `sql:"default:CURRENT_TIMESTAMP;"json:"created_at"`
	UpdatedAt time.Time `sql:"type:timestamp;default:CURRENT_TIMESTAMP;"json:"updated_at"`
	DeletedAt *time.Time `sql:"index"json:"deleted_at"`
}
// postgres不支持直接更新 需要触发器
