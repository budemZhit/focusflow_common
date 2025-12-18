package models

type SyncLinks struct {
	ID               int64  `db:"id" json:"id"`
	SourceCalendarID int64  `db:"source_calendar_id" json:"source_calendar_id"`
	TargetCalendarID int64  `db:"target_calendar_id" json:"target_calendar_id"`
	Direction        string `db:"direction" json:"direction"`
	Status           string `db:"status" json:"status"`
	Priority         int    `db:"priority" json:"priority"`
}
