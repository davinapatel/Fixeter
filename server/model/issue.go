// The package model describes all the models/data structures that will be used in the application
//
// The Issue struct holds all the data related to a single Issue that has been logged. Most of these fields
// the user will fill in via a form on a web page
//
// Fields:
// - ID:          The Primary Key of the table, a unique identifier for each record that autoincrements
// - Category:    The Category of Issue i.e Graffiti, Potholes
// - Title:       Title of the Issue
// - Date:        The date of the issue to be reported
// - Description: A description of the logged issue
// - Latitude:    The latitude coordinate of the location that the issue is located
// - Longitude:   The longitude coordinate of the location that the issue is located
// - Address:     The address of the location that the issue is located
// - Image:       The path of the where the image uploaded with the issue is located
// - Status:      Status of the issue i.e Logged, In Progress or Closed
// - ResourceID:  The foreign key to the Resource table/struct (an Issue will have a resource allocated to it)
// - Comments:    Comments added to Issue to update the progress
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
	Comments    string    `json:"comments" gorm:"column:comments"`

	Resource Resource `gorm:"foreignKey:ResourceID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
}
