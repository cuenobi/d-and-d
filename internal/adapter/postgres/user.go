package postgres

import (
	"errors"

	"d-and-d/internal/model"

	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{
		db: db,
	}
}

func (u *User) HasUsername(username string) (bool, error) {
	var count int64
	u.db.Model(&model.User{}).
		Where("username = ?", username).
		Count(&count)
	return count > 0, nil
}

func (u *User) CreateUser(user *model.User) error {
	result := u.db.Create(user)
	return result.Error
}

func (u *User) GetUserByUsername(username string) (*model.User, error) {
	var user *model.User
	err := u.db.First(&user, "username = ?", username).Error
	return user, err
}

func (u *User) CreateCharacter(character *model.Character) error {
	result := u.db.Create(character)
	return result.Error
}

func (u *User) CreateQuest(quest *model.Quest) error {
	result := u.db.Create(quest)
	return result.Error
}

func (u *User) GetAllCharacter() ([]*model.Character, error) {
	var characters []*model.Character
	err := u.db.
		Where("deleted_at IS NULL").
		Preload("Class").
		Preload("Race").
		Find(&characters).Error
	return characters, err
}

func (u *User) GetAllQuest() ([]*model.Quest, error) {
	var quests []*model.Quest
	err := u.db.
		Where("deleted_at IS NULL").
		Preload("DifficultyLevels").
		Find(&quests).Error
	return quests, err
}

func (u *User) UpdateCharacter(userID string, character *model.Character, privacyUpdate bool) error {
	var existingCharacter model.Character
	if err := u.db.Where("id = ? AND user_id = ?", character.ID, userID).
		First(&existingCharacter).Error; err != nil {
		return err
	}

	if character.Name == "" {
		character.Name = existingCharacter.Name
	}
	if character.Description == "" {
		character.Description = existingCharacter.Description
	}
	if character.RaceID == "" {
		character.RaceID = existingCharacter.RaceID
	}
	if character.ClassID == "" {
		character.ClassID = existingCharacter.ClassID
	}
	if len(character.Images) > 0 {
		character.Images = append(character.Images, existingCharacter.Images...)
	}
	if privacyUpdate {
		err := u.db.Model(&existingCharacter).
			Update("private", character.Private).
			Error
		if err != nil {
			return err
		}
	}

	result := u.db.Model(&existingCharacter).
		Updates(character)
	return result.Error
}

func (u *User) UpdateQuest(userID string, quest *model.Quest, privacyUpdate bool) error {
	var existingQuest model.Quest
	if err := u.db.Where("id = ? AND user_id = ?", quest.ID, userID).
		First(&existingQuest).Error; err != nil {
		return err
	}

	if quest.Name == "" {
		quest.Name = existingQuest.Name
	}
	if quest.Description == "" {
		quest.Description = existingQuest.Description
	}
	if quest.DifficultyLevelID == "" {
		quest.DifficultyLevelID = existingQuest.DifficultyLevelID
	}
	if len(quest.Images) > 0 {
		quest.Images = append(quest.Images, existingQuest.Images...)
	}
	if privacyUpdate {
		err := u.db.Model(&existingQuest).
			Update("private", quest.Private).
			Error
		if err != nil {
			return err
		}
	}

	result := u.db.Model(&existingQuest).
		Updates(quest)
	return result.Error
}

func (u *User) DeleteCharacter(characterID, userID string) error {
	result := u.db.Where("id = ? AND user_id = ?", characterID, userID).
		Delete(&model.Character{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no character found or you do not have permission to delete")
	}

	return nil
}

func (u *User) DeleteQuest(questID, userID string) error {
	result := u.db.Where("id = ? AND user_id = ?", questID, userID).
		Delete(&model.Quest{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no quest found or you do not have permission to delete")
	}

	return nil
}
