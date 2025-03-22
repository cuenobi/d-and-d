package model

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Class struct {
	ID          string         `gorm:"primaryKey"`
	CreatedAt   *time.Time     `gorm:"autoCreateTime"`
	UpdatedAt   *time.Time     `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string
	Description string
	Character   []*Character `gorm:"many2many:race_character;"`
}

type Race struct {
	ID          string         `gorm:"primaryKey"`
	CreatedAt   *time.Time     `gorm:"autoCreateTime"`
	UpdatedAt   *time.Time     `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string
	Description string
	Character   []*Character `gorm:"many2many:race_character;"`
}

type DifficultyLevels struct {
	ID          string         `gorm:"primaryKey"`
	CreatedAt   *time.Time     `gorm:"autoCreateTime"`
	UpdatedAt   *time.Time     `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string
	Description string
	Quests      []*Quest `gorm:"many2many:difficultyLevels_quest;"`
}

func (c *Class) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == "" {
		c.ID = strings.ReplaceAll(uuid.New().String(), "-", "")
	}
	return
}

func (r *Race) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == "" {
		r.ID = strings.ReplaceAll(uuid.New().String(), "-", "")
	}
	return
}

func (d *DifficultyLevels) BeforeCreate(tx *gorm.DB) (err error) {
	if d.ID == "" {
		d.ID = strings.ReplaceAll(uuid.New().String(), "-", "")
	}
	return
}
