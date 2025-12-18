package models

type SyncAudit struct {
	ID         int64  `db:"id" json:"id"`
	SyncLinkID int64  `db:"sync_link_id" json:"sync_link_id"`
	EntityType string `db:"entity_type" json:"entity_type"` // e.g., "event", "note", "category"
	EntityID   int64  `db:"entity_id" json:"entity_id"`
	Action     string `db:"action" json:"action"`           // e.g., "create", "update", "delete", "conflict_resolved"
	ResolvedBy string `db:"resolved_by" json:"resolved_by"` // e.g., user who resolved a conflict
	OccurredAt string `db:"occurred_at" json:"occurred_at"` // timestamp of the action
}
