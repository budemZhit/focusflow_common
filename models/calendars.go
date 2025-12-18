package models

type Calendars struct {
	ID          int64  `db:"id" json:"id"`
	OwnerID     int64  `db:"owner_id" json:"owner_id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
}
