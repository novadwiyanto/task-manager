package models

type Category struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"type:varchar(255);not null"`
	Description string `json:"description" gorm:"type:text"`
	UserID      uint   `json:"user_id"`
	User        User   `json:"user" gorm:"foreignKey:UserID"`
	TeamID      *uint  `json:"team_id"`
	Team        *Team  `json:"team" gorm:"foreignKey:TeamID"`
	CreatedAt   string `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   string `json:"updated_at" gorm:"autoUpdateTime"`
}
