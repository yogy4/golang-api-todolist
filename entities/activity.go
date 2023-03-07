package entities

import "time"

type Activity struct {
	// gorm.Model
	ActivityID int    `json:"id" gorm:"column:activity_id;primaryKey"`
	Title      string `json:"title"`
	Email      string `json:"email" gorm:"unique; not null"`
	Todo       []Todo `gorm:"foreignKey:ActivityRefer"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
