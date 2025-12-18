package models

type CategoryColors struct {
	ID         int64  `db:"id" json:"id"`
	UserID     int64  `db:"user_id" json:"user_id"`
	CategoryID int64  `db:"category_id" json:"category_id"`
	Color      string `db:"color" json:"color"`
}
