package port

import (
	"mime/multipart"

	"d-and-d/internal/constant"
	"d-and-d/internal/dto"
	"d-and-d/internal/model"

	"github.com/gofiber/fiber/v2"
)

type PublicService interface {
	GetPublicCharacter() ([]dto.PublicCharacterResponse, error)
	GetPublicQuest() ([]dto.PublicQuestResponse, error)
}

type UserService interface {
	CreateUser(ctx *fiber.Ctx, user *model.User) error
	CreateAdmin(ctx *fiber.Ctx, user *model.User) error
	Authentication(ctx *fiber.Ctx, username, password string) (string, uint, error)
	CreateCharacter(ctx *fiber.Ctx, character *model.Character, images []*multipart.FileHeader) error
	CreateQuest(ctx *fiber.Ctx, quest *model.Quest, images []*multipart.FileHeader) error
	SaveImage(ctx *fiber.Ctx, files []*multipart.FileHeader, userID string, imageType constant.ImageStoragePath) ([]string, error)
	RemoveFiles(targets []string) error
	GetAllCharacter(ctx *fiber.Ctx) ([]dto.CharacterResponse, error)
	GetAllQuest(ctx *fiber.Ctx) ([]dto.QuestResponse, error)
	UpdateCharacter(ctx *fiber.Ctx, character *model.Character, files []*multipart.FileHeader, privacyUpdate bool) error
	UpdateQuest(ctx *fiber.Ctx, quest *model.Quest, images []*multipart.FileHeader, privacyUpdate bool) error
	DeleteCharacter(ctx *fiber.Ctx, characterID string) error
	DeleteQuest(ctx *fiber.Ctx, questID string) error

	// TODO
	GetRace(ctx *fiber.Ctx) ([]*model.Race, error)
	GetClass(ctx *fiber.Ctx) ([]*model.Class, error)
	GetDifficultyLevel(ctx *fiber.Ctx) ([]*model.DifficultyLevels, error)
}

type AdminService interface {
	CreateClass(class *model.Class) error                     // TODO: Implement this
	CreateRace(race *model.Race) error                        // TODO: Implement this
	CreateDifficultyLevel(diff *model.DifficultyLevels) error // TODO: Implement this

	UpdateClass(class *model.Class) error                     // TODO: Implement this
	UpdateRace(race *model.Race) error                        // TODO: Implement this
	UpdateDifficultyLevel(diff *model.DifficultyLevels) error // TODO: Implement this

	DeleteClass(id uint) error           // TODO: Implement this
	DeleteRace(id uint) error            // TODO: Implement this
	DeleteDifficultyLevel(id uint) error // TODO: Implement this
}
