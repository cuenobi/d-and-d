package service

import (
	"errors"
	"fmt"
	"log/slog"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"d-and-d/config"
	"d-and-d/internal/constant"
	"d-and-d/internal/dto"
	"d-and-d/internal/model"
	"d-and-d/internal/port"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	logger      *slog.Logger
	userRepo    port.UserRepository
	optionsRepo port.OptionsRepository
	jwt         port.JWT
	cfg         *config.Config
}

func NewUserService(userRepo port.UserRepository, optionRepo port.OptionsRepository, jwt port.JWT, logger *slog.Logger, cfg *config.Config) *UserService {
	return &UserService{
		userRepo:    userRepo,
		optionsRepo: optionRepo,
		jwt:         jwt,
		logger:      logger,
		cfg:         cfg,
	}
}

func (u *UserService) CreateUser(ctx *fiber.Ctx, user *model.User) error {
	if constant.Role(user.Role) != constant.User {
		return fmt.Errorf("invalid role")
	}

	// Check if the username already exists in the system
	usernameExist, err := u.userRepo.HasUsername(user.Username)
	if err != nil {
		return err
	}

	if usernameExist {
		return fmt.Errorf("username already exist")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Set the hashed password to the user struct
	user.Password = string(hashedPassword)

	// Create the new user
	err = u.userRepo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) CreateAdmin(ctx *fiber.Ctx, user *model.User) error {
	if constant.Role(user.Role) != constant.Admin {
		return fmt.Errorf("invalid role")
	}

	// Check if the username already exists in the system
	usernameExist, err := u.userRepo.HasUsername(user.Username)
	if err != nil {
		return err
	}

	if usernameExist {
		return fmt.Errorf("username already exist")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Set the hashed password to the user struct
	user.Password = string(hashedPassword)

	// Create the new user
	err = u.userRepo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) Authentication(ctx *fiber.Ctx, username, password string) (string, uint, error) {
	// Retrieve the user from the database based on the username
	user, err := u.userRepo.GetUserByUsername(username)
	if err != nil {
		return "", 0, err
	}

	// Compare the provided password with the hashed password in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		if err.Error() == "crypto/bcrypt: hashedPassword is not the hash of the given password" {
			return "", 0, fmt.Errorf("invalid credential")
		}
		return "", 0, err
	}

	// Generate a JWT token for the user
	token := u.jwt.Generate(user.Username, user.Role)

	return token, user.Role, nil
}

func (u *UserService) GetRace(ctx *fiber.Ctx) ([]dto.RaceResponse, error) {
	resp := make([]dto.RaceResponse, 0)
	data, err := u.optionsRepo.GetAllRace()
	if err != nil {
		return nil, err
	}
	for _, d := range data {
		resp = append(resp, dto.RaceResponse{
			ID:          d.ID,
			Name:        d.Name,
			Description: d.Description,
		})
	}
	return resp, nil
}

func (u *UserService) GetClass(ctx *fiber.Ctx) ([]dto.ClassResponse, error) {
	resp := make([]dto.ClassResponse, 0)
	data, err := u.optionsRepo.GetAllClass()
	if err != nil {
		return nil, err
	}
	for _, d := range data {
		resp = append(resp, dto.ClassResponse{
			ID:          d.ID,
			Name:        d.Name,
			Description: d.Description,
		})
	}
	return resp, nil
}

func (u *UserService) GetDifficultyLevel(ctx *fiber.Ctx) ([]dto.DifficultyLevelResponse, error) {
	resp := make([]dto.DifficultyLevelResponse, 0)
	data, err := u.optionsRepo.GetAllDifficultyLevel()
	if err != nil {
		return nil, err
	}
	for _, d := range data {
		resp = append(resp, dto.DifficultyLevelResponse{
			ID:          d.ID,
			Name:        d.Name,
			Description: d.Description,
		})
	}
	return resp, nil
}

func (u *UserService) GetAllCharacter(ctx *fiber.Ctx) ([]dto.CharacterResponse, error) {
	resp := make([]dto.CharacterResponse, 0)
	characters, err := u.userRepo.GetAllCharacter()
	if err != nil {
		return nil, err
	}
	for _, character := range characters {
		resp = append(resp, dto.CharacterResponse{
			ID:          character.ID,
			Name:        character.Name,
			Description: character.Description,
			ClassName:   character.Class.Name,
			RaceName:    character.Race.Name,
			Images:      character.Images,
		})
	}
	return resp, nil
}

func (u *UserService) GetAllQuest(ctx *fiber.Ctx) ([]dto.QuestResponse, error) {
	resp := make([]dto.QuestResponse, 0)
	quests, err := u.userRepo.GetAllQuest()
	if err != nil {
		return nil, err
	}
	for _, quest := range quests {
		resp = append(resp, dto.QuestResponse{
			ID:              quest.ID,
			Name:            quest.Name,
			Description:     quest.Description,
			Difficulty:      quest.DifficultyLevels.Name,
			DiffDescription: quest.DifficultyLevels.Description,
			Images:          quest.Images,
		})
	}
	return resp, nil
}

func (u *UserService) CreateCharacter(ctx *fiber.Ctx, character *model.Character, files []*multipart.FileHeader) error {
	class, err := u.optionsRepo.GetClassByID(character.ClassID)
	if err != nil {
		return err
	}

	race, err := u.optionsRepo.GetRaceByID(character.RaceID)
	if err != nil {
		return err
	}

	username, ok := ctx.Locals("username").(string)
	if !ok {
		return errors.New("username is not a valid string")
	}

	user, err := u.userRepo.GetUserByUsername(username)
	if err != nil {
		return err
	}

	var imagePaths []string
	if len(files) > 0 {
		imagePaths, err = u.SaveImage(ctx, files, user.ID, constant.CharacterPath)
		if err != nil {
			u.logger.Error(err.Error())
			return err
		}
	}

	character = &model.Character{
		Name:        character.Name,
		Description: character.Description,
		RaceID:      race.ID,
		ClassID:     class.ID,
		Images:      imagePaths,
		Private:     character.Private,
		UserID:      user.ID,
	}

	if err := u.userRepo.CreateCharacter(character); err != nil {
		u.logger.Error(err.Error())
		_ = u.RemoveFiles(imagePaths)
		return err
	}

	return nil
}

func (u *UserService) CreateQuest(ctx *fiber.Ctx, quest *model.Quest, files []*multipart.FileHeader) error {
	username, ok := ctx.Locals("username").(string)
	if !ok {
		return errors.New("username is not a valid string")
	}

	user, err := u.userRepo.GetUserByUsername(username)
	if err != nil {
		return err
	}

	diffLevel, err := u.optionsRepo.GetDifficultyLevelById(quest.DifficultyLevelID)
	if err != nil {
		u.logger.Error(err.Error())
		return err
	}

	var imagePaths []string
	if len(files) > 0 {
		imagePaths, err = u.SaveImage(ctx, files, user.ID, constant.QuestPath)
		if err != nil {
			u.logger.Error(err.Error())
			return err
		}
	}

	quest = &model.Quest{
		Name:              quest.Name,
		Description:       quest.Description,
		Images:            imagePaths,
		UserID:            user.ID,
		DifficultyLevelID: diffLevel.ID,
		DifficultyLevels:  diffLevel,
		Private:           quest.Private,
	}

	if err := u.userRepo.CreateQuest(quest); err != nil {
		u.logger.Error(err.Error())
		_ = u.RemoveFiles(imagePaths)
		return err
	}

	return nil
}

func (u *UserService) UpdateCharacter(ctx *fiber.Ctx, character *model.Character, files []*multipart.FileHeader, privacyUpdate bool) error {
	username, ok := ctx.Locals("username").(string)
	if !ok {
		return errors.New("username is not a valid string")
	}

	user, err := u.userRepo.GetUserByUsername(username)
	if err != nil {
		return err
	}

	classID := ""
	if character.ClassID != "" {
		class, err := u.optionsRepo.GetClassByID(character.ClassID)
		if err != nil {
			u.logger.Error(err.Error())
			return err
		}
		classID = class.ID
	}

	raceID := ""
	if character.RaceID != "" {
		race, err := u.optionsRepo.GetRaceByID(character.RaceID)
		if err != nil {
			u.logger.Error(err.Error())
			return err
		}
		raceID = race.ID
	}

	var imagePaths []string
	if len(files) > 0 {
		imagePaths, err = u.SaveImage(ctx, files, user.ID, constant.QuestPath)
		if err != nil {
			u.logger.Error(err.Error())
			return err
		}
	}

	character = &model.Character{
		ID:          character.ID,
		Name:        character.Name,
		Description: character.Description,
		Images:      imagePaths,
		UserID:      user.ID,
		Private:     character.Private,
		RaceID:      raceID,
		ClassID:     classID,
	}

	if err := u.userRepo.UpdateCharacter(user.ID, character, privacyUpdate); err != nil {
		u.logger.Error(err.Error())
		_ = u.RemoveFiles(imagePaths)
		return err
	}

	return nil
}

func (u *UserService) UpdateQuest(ctx *fiber.Ctx, quest *model.Quest, files []*multipart.FileHeader, privacyUpdate bool) error {
	username, ok := ctx.Locals("username").(string)
	if !ok {
		return errors.New("username is not a valid string")
	}

	user, err := u.userRepo.GetUserByUsername(username)
	if err != nil {
		return err
	}

	diffLevelID := ""
	if quest.DifficultyLevelID != "" {
		diffLevel, err := u.optionsRepo.GetDifficultyLevelById(quest.DifficultyLevelID)
		if err != nil {
			u.logger.Error(err.Error())
			return err
		}
		diffLevelID = diffLevel.ID
	}

	var imagePaths []string
	if len(files) > 0 {
		imagePaths, err = u.SaveImage(ctx, files, user.ID, constant.QuestPath)
		if err != nil {
			u.logger.Error(err.Error())
			return err
		}
	}

	quest = &model.Quest{
		ID:                quest.ID,
		Name:              quest.Name,
		Description:       quest.Description,
		Images:            imagePaths,
		UserID:            user.ID,
		Private:           quest.Private,
		DifficultyLevelID: diffLevelID,
	}

	if err := u.userRepo.UpdateQuest(user.ID, quest, privacyUpdate); err != nil {
		u.logger.Error(err.Error())
		_ = u.RemoveFiles(imagePaths)
		return err
	}

	return nil
}

func (u *UserService) DeleteCharacter(ctx *fiber.Ctx, characterID string) error {
	username, ok := ctx.Locals("username").(string)
	if !ok {
		return errors.New("username is not a valid string")
	}

	user, err := u.userRepo.GetUserByUsername(username)
	if err != nil {
		return err
	}

	return u.userRepo.DeleteCharacter(characterID, user.ID)
}

func (u *UserService) DeleteQuest(ctx *fiber.Ctx, questID string) error {
	username, ok := ctx.Locals("username").(string)
	if !ok {
		return errors.New("username is not a valid string")
	}

	user, err := u.userRepo.GetUserByUsername(username)
	if err != nil {
		return err
	}

	return u.userRepo.DeleteQuest(questID, user.ID)
}

func (u *UserService) SaveImage(ctx *fiber.Ctx, files []*multipart.FileHeader, userID string, imageType constant.ImageStoragePath) ([]string, error) {
	uploadDir := fmt.Sprintf("%s/%s/%s", u.cfg.ServerConfig.FilePath.Image, userID, imageType)
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return nil, err
	}

	var imagePaths []string
	for _, file := range files {
		filename := fmt.Sprintf("%d-%s", time.Now().UnixNano(), file.Filename)
		savePath := filepath.Join(uploadDir, filename)

		if err := ctx.SaveFile(file, savePath); err != nil {
			return nil, err
		}
		savePath = strings.TrimPrefix(savePath, u.cfg.ServerConfig.FilePath.Image)
		imagePaths = append(imagePaths, savePath)
	}

	return imagePaths, nil
}

func (u *UserService) RemoveFiles(targets []string) error {
	for _, file := range targets {
		err := os.Remove(file)
		if err != nil {
			u.logger.Error(err.Error())
		}
	}
	return nil
}
