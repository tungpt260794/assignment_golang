package models

import "time"

// Photo ...
type Photo struct {
	ID          int        `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	AccountID   int        `gorm:"NOT NULL" json:"account_id"`
	GalleryID   int        `gorm:"NOT NULL" json:"gallery_id"`
	Name        string     `gorm:"type:TEXT;NOT NULL" json:"name"`
	Description string     `gorm:"type:TEXT;" json:"description"`
	Size        int64      `gorm:"type:DOUBLE;NOT NULL" json:"size"`
	Path        string     `gorm:"type:TEXT;NOT NULL"  json:"path"`
	W1920Path   string     `gorm:"type:VARCHAR(265);NOT NULL" json:"w1920_path"`
	W1600Path   string     `gorm:"type:VARCHAR(265);NOT NULL" json:"w1600_path"`
	W1280Path   string     `gorm:"type:VARCHAR(265);NOT NULL" json:"w1280_path"`
	W1024Path   string     `gorm:"type:VARCHAR(265);NOT NULL" json:"w1024_path"`
	W800Path    string     `gorm:"type:VARCHAR(265);NOT NULL" json:"w800_path"`
	W256Path    string     `gorm:"type:VARCHAR(265);NOT NULL" json:"w256_path"`
	CreatedAt   time.Time  `gorm:"NOT NULL" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"NOT NULL" json:"updated_at"`
	Reactions   []Reaction `json:"reactions,omitempty"`
}
