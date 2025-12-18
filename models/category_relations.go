package models

type CategoryRelations struct {
	ID                      int64  `db:"id" json:"id"`
	UserID                  int64  `db:"user_id" json:"user_id"`
	PrimaryCategoryID       int64  `db:"primary_category_id" json:"primary_category_id"`
	SecondaryCategoryID     int64  `db:"secondary_category_id" json:"secondary_category_id"`
	RepeatPattern           string `db:"repeat_pattern" json:"repeat_pattern"`
	StartDate               string `db:"start_date" json:"start_date"`
	CheckProductionCalendar bool   `db:"check_production_calendar" json:"check_production_calendar"`
}
