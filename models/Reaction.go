package models

import "time"

// Reaction ...
type Reaction struct {
	AccountID int       `gorm:"PRIMARY_KEY;AUTO_INCREMENT:false;NOT NULL" json:"account_id"`
	PhotoID   int       `gorm:"PRIMARY_KEY;AUTO_INCREMENT:false;NOT NULL" json:"photo_id"`
	CreatedAt time.Time `gorm:"NOT NULL" json:"created_at"`
	UpdatedAt time.Time `gorm:"NOT NULL" json:"updated_at"`
}
