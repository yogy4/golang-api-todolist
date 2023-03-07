package entities

import "time"

type Todo struct {
	// gorm.Model
	ID            int `json:"id" gorm:"column:todo_id; primaryKey"`
	ActivityRefer int `json:"activity_group_id" gorm:"column:activity_group_id"`

	Title      string `json:"title"`
	IsActivate bool   `json:"is_activate" gorm:"default:true"`
	Priority   string `json:"priority"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
