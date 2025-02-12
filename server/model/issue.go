package model

import (
	"time"
)

// Can map struct fields into JSON Fields

type Issue struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Category    string    `json:"category" gorm:"not null;column:category;size:255"`
	Title       string    `json:"title" gorm:"not null;column:title;size:255"`
	Date        time.Time `json:"date" gorm:"not null;column:date;default:CURRENT_TIMESTAMP"`
	Description string    `json:"description" gorm:"not null;column:description;type:text"`
	Image       string    `json:"image" gorm:"null;column:image;size:255"`
}
