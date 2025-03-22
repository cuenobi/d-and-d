package model

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          string         `gorm:"primaryKey"`
	CreatedAt   *time.Time     `gorm:"autoCreateTime"`
	UpdatedAt   *time.Time     `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Username    string         `gorm:"unique"`
	Password    string
	Name        string
	Role        uint
	CharacterID string     `gorm:"index"`                  // Foreign Key สำหรับ Character
	Character   *Character `gorm:"foreignKey:CharacterID"` // Foreign Key เชื่อมไปที่ Character
	Quests      []*Quest   `gorm:"many2many:user_quests;"`
}

type UserQuest struct {
	UserID  string `gorm:"primaryKey"`
	QuestID string `gorm:"primaryKey"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = strings.ReplaceAll(uuid.New().String(), "-", "")
	}
	return
}
