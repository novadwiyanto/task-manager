package models

import (
	"task-manager/pkg/enum"
)

type Notification struct {
	ID      int                   `json:"id" gorm:"primaryKey"`
	UserID  int                   `json:"user_id" gorm:"not null"`
	TaskID  int                   `json:"task_id" gorm:"not null"`
	Type    enum.NotificationType `json:"type" gorm:"type:ENUM('assigned', 'updated', 'deleted'); default:'assigned';not null"`
	Message string                `json:"message" gorm:"type:text;not null"`
	IsRead  bool                  `json:"is_read" gorm:"default:false"`
}
