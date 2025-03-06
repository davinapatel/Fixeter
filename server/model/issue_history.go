package model

type IssueHistory struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Category string `json:"category" gorm:"not null;column:category;size:255"`
	Count    int    `json:"count" gorm:"not null;column:count"`
}
