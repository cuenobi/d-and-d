package model

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Quest struct {
	ID                string         `gorm:"primaryKey"`
	CreatedAt         *time.Time     `gorm:"autoCreateTime"`
	UpdatedAt         *time.Time     `gorm:"autoUpdateTime"`
	DeletedAt         gorm.DeletedAt `gorm:"index"`
	Name              string
	Description       string
	DifficultyLevelID string            `gorm:"index"`
	DifficultyLevels  *DifficultyLevels `gorm:"foreignKey:DifficultyLevelID"`
	Images            pq.StringArray    `gorm:"type:text[]"`
	Private           bool
	UserID            string
}

func (q *Quest) BeforeCreate(tx *gorm.DB) (err error) {
	if q.ID == "" {
		q.ID = strings.ReplaceAll(uuid.New().String(), "-", "")
	}
	return
}
