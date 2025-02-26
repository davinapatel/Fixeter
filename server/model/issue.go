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
	Latitude    float64   `json:"latitude" gorm:"not null;column:latitude"`
	Longitude   float64   `json:"longitude" gorm:"not null;column:longitude"`
	Address     string    `json:"address" gorm:"not null;column:address;type:text"`
	Image       string    `json:"image" gorm:"null;column:image;size:255"`
	Status      string    `json:"status" gorm:"not null;column:status;size:255"`
	ResourceID  uint      `json:"resourceId" gorm:"not null;column:resourceId"`

	Resource Resource `gorm:"foreignKey:ResourceID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
}
