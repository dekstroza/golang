package models

//ApplicationUser describes single entry in that table
type ApplicationUser struct {
	ID        uint64 `gorm:"index:idx_app_user_id"`
	Firstname string
	Lastname  string
}
