// The package model describes all the models/data structures that will be used in the application
//
// The IssueHistory struct holds the data for the number of Issues reported per Category over a period
// of time.
//
// Fields:
// - ID:        The Primary Key of the table, a unique identifier for each record that autoincrements
// - Category:  The Category of Issue i.e Graffiti, Potholes
// - Count:     The number of times an Issue has been reported for that Category, this field will be updated
//              every time an Issue of that Category type has been logged

package model

type IssueHistory struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"` // Can map struct fields into JSON fields
	Category string `json:"category" gorm:"not null;column:category;size:255"`
	Count    int    `json:"count" gorm:"not null;column:count"`
}
