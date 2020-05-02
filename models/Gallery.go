package models

import "time"

// Gallery ...
type Gallery struct {
	ID         int       `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	AccountID  int       `gorm:"NOT NULL" json:"account_id"`
	Name       string    `gorm:"type:TEXT;NOT NULL" json:"name"`
	Brief      string    `gorm:"type:TEXT;NOT NULL" json:"brief"`
	Visibility string    `gorm:"type:VARCHAR(64);default:'private'" json:"visibility"`
	CreatedAt  time.Time `gorm:"NOT NULL" json:"created_at"`
	UpdatedAt  time.Time `gorm:"NOT NULL" json:"updated_at"`
	Photos     []Photo   `json:"photos,omitempty"`
}
