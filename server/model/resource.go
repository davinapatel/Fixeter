// The package model describes all the models/data structures that will be used in the application
//
// The Resource struct holds the data for the resources that can be allocated to an Issue
//
// Fields:
// - ID:          The Primary Key of the table, a unique identifier for each record that autoincrements
// - Type:        The type of resource i.e Binmen to visit location, Graffiti removal
// - Department:  The department which the type of resource belongs to i.e Waste Removal

package model

// Can map struct fields into JSON Fields

type Resource struct {
	ID         uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	Type       string  `json:"type" gorm:"not null;column:type;size:255"`
	Department string  `json:"department" gorm:"not null;column:department;size:255"`
	Issues     []Issue `gorm:"foreignKey:ResourceID"`
}
