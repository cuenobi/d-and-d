package port

import "d-and-d/internal/model"

type PublicRepository interface {
	GetPublicQuest() ([]*model.Quest, error)
	GetPublicCharacter() ([]*model.Character, error)
}

type UserRepository interface {
	HasUsername(username string) (bool, error)
	GetUserByUsername(username string) (*model.User, error)
	CreateUser(user *model.User) error
	CreateCharacter(character *model.Character) error
	CreateQuest(quest *model.Quest) error
	GetAllCharacter() ([]*model.Character, error)
	GetAllQuest() ([]*model.Quest, error)
	UpdateCharacter(userID string, character *model.Character, privacyUpdate bool) error
	UpdateQuest(userID string, quest *model.Quest, privacyUpdate bool) error
	DeleteCharacter(characterID, userID string) error
	DeleteQuest(questID, userID string) error
}

type OptionsRepository interface {
	GetAllRace() ([]*model.Race, error)
	GetRaceByID(id string) (*model.Race, error)
	GetAllClass() ([]*model.Class, error)
	GetClassByID(id string) (*model.Class, error)
	GetAllDifficultyLevel() ([]*model.DifficultyLevels, error)
	GetDifficultyLevelById(id string) (*model.DifficultyLevels, error)
}
