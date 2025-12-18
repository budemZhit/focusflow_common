package models

type CalendarAccess struct {
	ID                  int64   `db:"id" json:"id"`
	CalendarID          int64   `db:"calendar_id" json:"calendar_id"`
	UserID              int64   `db:"user_id" json:"user_id"`
	CanAddCategories    bool    `db:"can_add_categories" json:"can_add_categories"`
	CanAddNotes         bool    `db:"can_add_notes" json:"can_add_notes"`
	CanEditDays         bool    `db:"can_edit_days" json:"can_edit_days"`
	CanEditCategories   bool    `db:"can_edit_categories" json:"can_edit_categories"`
	CanEditColors       bool    `db:"can_edit_colors" json:"can_edit_colors"`
	CanViewAll          bool    `db:"can_view_all" json:"can_view_all"`
	CanViewGeneralNotes bool    `db:"can_view_general_notes" json:"can_view_general_notes"`
	VisibleCategories   []int64 `db:"visible_categories" json:"visible_categories"`
}
