package model

// Can map struct fields into JSON Fields

type Resource struct {
	ID         uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	Type       string  `json:"type" gorm:"not null;column:type;size:255"`
	Department string  `json:"department" gorm:"not null;column:department;size:255"`
	Issues     []Issue `gorm:"foreignKey:ResourceID"`
}
