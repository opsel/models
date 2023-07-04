package models

import "gorm.io/plugin/soft_delete"

type Users struct {
	ID        uint64                `gorm:"column:id;primaryKey;autoIncrement"`
	FirstName string                `gorm:"column:name_first"`
	LastName  string                `gorm:"column:name_last"`
	Email     string                `gorm:"column:email;unique"`
	Password  string                `gorm:"column:password"`
	CreatedAt int64                 `gorm:"column:created_at;autoCreateTime;default:null"`
	UpdatedAt int64                 `gorm:"column:updated_at;autoUpdateTime;default:null"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;default:null;index"`
}

func (Users) TableName() string {
	return "users"
}
