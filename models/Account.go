package models

import (
	"time"
)

// Account ...
type Account struct {
	ID        int       `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	Email     string    `gorm:"type:VARCHAR(128);NOT NULL" json:"email"`
	Password  string    `gorm:"type:VARCHAR(64);NOT NULL" json:"password"`
	Name      string    `gorm:"type:VARCHAR(256);" json:"name"`
	Phone     string    `gorm:"type:VARCHAR(20);" json:"phone"`
	Address   string    `gorm:"type:TEXT;" json:"address"`
	Avatar    string    `gorm:"type:VARCHAR(256);" json:"avatar"`
	CreatedAt time.Time `gorm:"NOT NULL" json:"created_at"`
	UpdatedAt time.Time `gorm:"NOT NULL" json:"updated_at"`
}
