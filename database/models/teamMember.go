package models

import "task-manager/pkg/enum"

type TeamMember struct {
	ID     uint      `json:"id" gorm:"primaryKey"`
	TeamID int       `json:"team_id" gorm:"not null"`
	Team   Team      `json:"team" gorm:"foreignKey:TeamID"`
	Role   enum.Role `json:"role" gorm:"type:ENUM('admin', 'member', 'guest');not null;default:'guest'"`
	UserID int       `json:"user_id" gorm:"not null"`
	User   User      `json:"user" gorm:"foreignKey:UserID"`
}
