package postgres

import (
	"errors"
	"time"

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
	var races []*model.Race
	err := o.db.Model(&model.Race{}).
		Where("deleted_at IS NULL").
		Find(&races).Error
	if err != nil {
		return nil, err
	}
	return races, nil
}

func (o *Option) GetAllClass() ([]*model.Class, error) {
	var classes []*model.Class
	err := o.db.Model(&model.Class{}).
		Where("deleted_at IS NULL").
		Find(&classes).Error
	if err != nil {
		return nil, err
	}
	return classes, nil
}

func (o *Option) GetAllDifficultyLevel() ([]*model.DifficultyLevels, error) {
	var diffs []*model.DifficultyLevels
	err := o.db.Model(&model.DifficultyLevels{}).
		Where("deleted_at IS NULL").
		Find(&diffs).Error
	if err != nil {
		return nil, err
	}
	return diffs, nil
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

func (o *Option) CreateRace(race *model.Race) error {
	return o.db.Create(race).Error
}

func (o *Option) CreateClass(class *model.Class) error {
	return o.db.Create(class).Error
}

func (o *Option) CreateDifficultyLevel(diffLv *model.DifficultyLevels) error {
	return o.db.Create(diffLv).Error
}

func (o *Option) UpdateRace(race *model.Race) error {
	result := o.db.Model(race).
		Where("id = ?", race.ID).
		Updates(race)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("race_id not found")
	}

	return nil
}

func (o *Option) UpdateClass(class *model.Class) error {
	result := o.db.Model(class).
		Where("id = ?", class.ID).
		Updates(class)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("class_id not found")
	}

	return nil
}

func (o *Option) UpdateDifficultyLevel(diffLv *model.DifficultyLevels) error {
	result := o.db.Model(diffLv).
		Where("id = ?", diffLv.ID).
		Updates(diffLv)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("difficulty_id not found")
	}

	return nil
}

func (o *Option) DeleteClass(id string) error {
	tx := o.db.Begin()

	var exitingCharacter *model.Character
	if err := tx.Where("class_id = ?", id).
		First(&exitingCharacter).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&model.Character{}).
		Where("class_id = ?", id).
		Update("deleted_at", gorm.DeletedAt{
			Time:  time.Now(),
			Valid: true,
		}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("id = ?", id).Delete(&model.Class{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (o *Option) DeleteRace(id string) error {
	tx := o.db.Begin()

	var exitingCharacter *model.Character
	if err := tx.Where("race_id = ?", id).
		First(&exitingCharacter).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&model.Character{}).
		Where("race_id = ?", id).
		Update("deleted_at", gorm.DeletedAt{
			Time:  time.Now(),
			Valid: true,
		}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("id = ?", id).Delete(&model.Race{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (o *Option) DeleteDifficultyLevel(id string) error {
	tx := o.db.Begin()

	var exitingQuest *model.Quest
	if err := tx.Where("difficulty_level_id = ?", id).
		First(&exitingQuest).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&model.Quest{}).
		Where("difficulty_level_id = ?", id).
		Update("deleted_at", gorm.DeletedAt{
			Time:  time.Now(),
			Valid: true,
		}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("id = ?", id).Delete(&model.DifficultyLevels{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
