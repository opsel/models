package models

import "gorm.io/plugin/soft_delete"

type Agent struct {
	ID        uint64                `gorm:"column:id;primaryKey"`
	UUID      string                `gorm:"column:uuid"`
	CreatedAt int64                 `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt int64                 `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;default:null"`
}

func (Agent) TableName() string {
	return "agent"
}
