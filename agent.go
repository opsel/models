package models

import "gorm.io/plugin/soft_delete"

type Agent struct {
	ID        uint64                `gorm:"column:id;primaryKey;autoIncrement"`
	UUID      string                `gorm:"column:uuid;unique"`
	CreatedAt int64                 `gorm:"column:created_at;autoCreateTime;default:null"`
	UpdatedAt int64                 `gorm:"column:updated_at;autoUpdateTime;default:null"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;default:null;index"`
}

func (Agent) TableName() string {
	return "agent"
}
