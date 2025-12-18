package models

type SyncFilters struct {
	ID               int64  `db:"id" json:"id"`
	SyncLinkID       int64  `db:"sync_link_id" json:"sync_link_id"`
	ScopeEvents      bool   `db:"scope_events" json:"scope_events"`
	ScopeNotes       bool   `db:"scope_notes" json:"scope_notes"`
	ViewGeneralNotes bool   `db:"view_general_notes" json:"view_general_notes"`
	CategoriesJSON   string `db:"categories_json" json:"categories_json"`
	DateRangeJSON    string `db:"date_range_json" json:"date_range_json"`
	TagsJSON         string `db:"tags_json" json:"tags_json"`
}
