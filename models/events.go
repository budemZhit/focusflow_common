package models

type Events struct {
	ID           int64  `db:"id" json:"id"`
	CalendarID   int64  `db:"calendar_id" json:"calendar_id"`
	Date         string `db:"date" json:"date"`
	CategoryID   int64  `db:"category_id" json:"category_id"`
	Title        string `db:"title" json:"title"`
	Comment      string `db:"comment" json:"comment"`
	RepeatRule   string `db:"repeat_rule" json:"repeat_rule"`
	Provenance   string `db:"provenance" json:"provenance"`
	SourceLinkID *int64 `db:"source_link_id" json:"source_link_id,omitempty"`
}
