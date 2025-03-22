package postgres

import (
	"d-and-d/internal/model"

	"gorm.io/gorm"
)

type Public struct {
	db *gorm.DB
}

func NewPublic(db *gorm.DB) *Public {
	return &Public{
		db: db,
	}
}

func (p *Public) GetPublicQuest() ([]*model.Quest, error) {
	var quests []*model.Quest
	err := p.db.Model(&model.Quest{}).
		Where("private = false").
		Where("deleted_at IS NULL").
		Preload("DifficultyLevels").
		Find(&quests).Error
	if err != nil {
		return nil, err
	}
	return quests, nil
}

func (p *Public) GetPublicCharacter() ([]*model.Character, error) {
	var characters []*model.Character
	err := p.db.Model(&model.Character{}).
		Where("private = false").
		Where("deleted_at IS NULL").
		Preload("Class").
		Preload("Race").
		Find(&characters).Error
	if err != nil {
		return nil, err
	}
	return characters, nil
}
