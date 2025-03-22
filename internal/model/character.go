package model

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Character struct {
	ID          string         `gorm:"primaryKey"`
	CreatedAt   *time.Time     `gorm:"autoCreateTime"`
	UpdatedAt   *time.Time     `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string
	Description string
	RaceID      string         `gorm:"index"`
	Race        *Race          `gorm:"foreignKey:RaceID"`
	ClassID     string         `gorm:"index"`
	Class       *Class         `gorm:"foreignKey:ClassID"`
	Images      pq.StringArray `gorm:"type:text[]"`
	Private     bool
	UserID      string
}

func (c *Character) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == "" {
		c.ID = strings.ReplaceAll(uuid.New().String(), "-", "")
	}
	return
}
