package models

type Categories struct {
	ID            int64  `db:"id" json:"id"`
	UserID        *int64 `db:"user_id" json:"user_id"` // nullable для системных категорий
	Name          string `db:"name" json:"name"`
	IsStandard    bool   `db:"is_standard" json:"is_standard"`
	RepeatRule    string `db:"repeat_rule" json:"repeat_rule"`       // daily, weekly, monthly, custom
	RepeatPattern string `db:"repeat_pattern" json:"repeat_pattern"` // JSON‑шаблон повторяемости
}
