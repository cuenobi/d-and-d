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
	GetRaceByID(id string) (*model.Race, error)
	GetClassByID(id string) (*model.Class, error)
	GetDifficultyLevelById(id string) (*model.DifficultyLevels, error)

	GetAllRace() ([]*model.Race, error)
	GetAllClass() ([]*model.Class, error)
	GetAllDifficultyLevel() ([]*model.DifficultyLevels, error)

	CreateRace(race *model.Race) error
	CreateClass(class *model.Class) error
	CreateDifficultyLevel(diffLevel *model.DifficultyLevels) error

	UpdateRace(race *model.Race) error
	UpdateClass(class *model.Class) error
	UpdateDifficultyLevel(diffLevel *model.DifficultyLevels) error

	DeleteRace(raceID string) error
	DeleteClass(classID string) error
	DeleteDifficultyLevel(diffLvID string) error
}
