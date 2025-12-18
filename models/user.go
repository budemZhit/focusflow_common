package models

type User struct {
	ID           int64  `db:"id" json:"id"`
	Name         string `db:"name" json:"name"`
	Email        string `db:"email" json:"email"`
	PasswordHash string `db:"password_hash" json:"password_hash"`
	AuthProvider string `db:"auth_provider" json:"auth_provider"`
	ProviderUID  string `db:"provider_uid" json:"provider_uid"`
	CreatedAt    int64  `db:"created_at" json:"created_at"`
}
