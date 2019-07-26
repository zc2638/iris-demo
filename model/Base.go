package model

import "time"

type AutoID struct {
	ID uint `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
}

type Timestamps struct {
	CreatedAt time.Time  `gorm:"type:timestamp;not null;" json:"created_at"`
	UpdatedAt time.Time  `gorm:"type:timestamp;not null;" json:"updated_at"`
	DeletedAt *time.Time `gorm:"type:timestamp;" json:"deleted_at"`
}
