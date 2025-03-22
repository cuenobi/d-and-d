package postgres

import (
	"d-and-d/internal/model"

	"gorm.io/gorm"
)

type Option struct {
	db *gorm.DB
}

func NewOption(db *gorm.DB) *Option {
	return &Option{
		db: db,
	}
}

func (o *Option) GetAllRace() ([]*model.Race, error) {
	return nil, nil
}

func (o *Option) GetRaceByID(id string) (*model.Race, error) {
	var race *model.Race
	err := o.db.Model(&model.Race{}).
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		First(&race).Error
	if err != nil {
		return nil, err
	}
	return race, nil
}

func (o *Option) GetAllClass() ([]*model.Class, error) {
	return nil, nil
}

func (o *Option) GetClassByID(id string) (*model.Class, error) {
	var class *model.Class
	err := o.db.Model(&model.Class{}).
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		First(&class).Error
	if err != nil {
		return nil, err
	}
	return class, nil
}

func (o *Option) GetDifficultyLevelById(id string) (*model.DifficultyLevels, error) {
	var difficulty *model.DifficultyLevels
	err := o.db.Model(&model.DifficultyLevels{}).
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		First(&difficulty).Error
	if err != nil {
		return nil, err
	}
	return difficulty, nil
}

func (o *Option) GetAllDifficultyLevel() ([]*model.DifficultyLevels, error) {
	return nil, nil
}
