package models

type Notes struct {
	ID         int64  `db:"id" json:"id"`
	CalendarID int64  `db:"calendar_id" json:"calendar_id"`
	Date       string `db:"date" json:"date"` // Format: YYYY-MM-DD
	Content    string `db:"content" json:"content"`
	CategoryID *int64 `db:"category_id" json:"category_id,omitempty"`
	IsGeneral  bool   `db:"is_general" json:"is_general"`
	Provenance string `db:"provenance" json:"provenance"`
}
