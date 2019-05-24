package model

import (
	"time"
)

type BaseModel struct {
	Id        uint64    `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	CreatedAt time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"-"`
}
