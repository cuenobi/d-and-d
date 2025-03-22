package handler

import (
	"errors"
	"log/slog"
	"strconv"
	"unicode/utf8"

	"d-and-d/config"
	"d-and-d/internal/dto"
	"d-and-d/internal/model"
	"d-and-d/internal/port"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type UserHandler struct {
	Fiber       *fiber.App
	UserService port.UserService
	JWT         port.JWT
	Validator   *validator.Validate
	Logger      *slog.Logger
	Config      *config.Config
}

type RegisterBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Role     uint   `json:"role" validate:"required"`
}

// RegisterHandler godoc
// @Summary Register a new user
// @Description Create a new user account
// @Tags User
// @Accept  json
// @Produce  json
// @Param request body RegisterBody true "User registration details"
// @Success 201 {object} dto.RegisterResponse
// @Failure 400 {object} ErrorResponse
// @Router /user/register [post]
func (u *UserHandler) RegisterHandler(ctx *fiber.Ctx) error {
	var input RegisterBody
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	if err := u.Validator.Struct(input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	user := &model.User{
		Username: input.Username,
		Password: input.Password,
		Name:     input.Name,
		Role:     input.Role,
	}

	if err := u.UserService.CreateUser(ctx, user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(dto.RegisterResponse{
		Message: "register success",
	})
}

type LoginBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// Login godoc
// @Summary User login
// @Description Authenticate user and return JWT token
// @Tags User
// @Accept  json
// @Produce  json
// @Param request body LoginBody true "User login details"
// @Success 200 {object} dto.RegisterResponse
// @Failure 400 {object} ErrorResponse
// @Router /login [post]
func (u *UserHandler) Login(ctx *fiber.Ctx) error {
	var input LoginBody
	if err := ctx.BodyParser(&input); err != nil {
		u.Logger.Error("error parsing body: ", "", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	if err := u.Validator.Struct(input); err != nil {
		u.Logger.Error("error validator ", "", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	token, role, err := u.UserService.Authentication(ctx, input.Username, input.Password)
	if err != nil {
		u.Logger.Error("error authentication", "", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.LoginResponse{
		Token:   token,
		Role:    role,
		Message: "login success",
	})
}

// RegisterHandler godoc
// @Security BearerAuth
// @Summary Get all quests
// @Description Get all quests
// @Tags Quest
// @Accept json
// @Produce json
// @Success 201 {object} dto.HandlerResponse
// @Failure 400 {object} dto.HandlerResponse
// @Router /user/quests [get]
func (u *UserHandler) GetAllQuest(ctx *fiber.Ctx) dto.HandlerResponse {
	quests, err := u.UserService.GetAllQuest(ctx)
	if err != nil {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      err,
		}
	}

	return dto.HandlerResponse{
		StatusCode: fiber.StatusOK,
		Message:    "Success",
		Data:       quests,
	}
}

// RegisterHandler godoc
// @Security BearerAuth
// @Summary Get all characters
// @Description Get all characters
// @Tags Character
// @Accept json
// @Produce json
// @Success 201 {object} dto.HandlerResponse
// @Failure 400 {object} dto.HandlerResponse
// @Router /user/characters [get]
func (u *UserHandler) GetAllCharacter(ctx *fiber.Ctx) dto.HandlerResponse {
	characters, err := u.UserService.GetAllCharacter(ctx)
	if err != nil {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      err,
		}
	}

	return dto.HandlerResponse{
		StatusCode: fiber.StatusOK,
		Message:    "Success",
		Data:       characters,
	}
}

// CreateQuest godoc
// @Summary Create a new quest
// @Description Allows a registered user to create a new quest with images
// @Tags Quest
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "Quest Name"
// @Param description formData string false "Quest Description (Max: 5000 digits)"
// @Param diff_id formData string true "Difficulty Level ID"
// @Param privacy formData boolean true "Quest Privacy (true = private, false = public)"
// @Param images formData file false "Quest Images (Max: 10 files)" collectionFormat=multi
// @Success 201 {object} dto.HandlerResponse
// @Failure 400 {object} dto.HandlerResponse
// @Router /user/quest [post]
// @Security BearerAuth
func (u *UserHandler) CreateQuest(ctx *fiber.Ctx) dto.HandlerResponse {
	name := ctx.FormValue("name")
	description := ctx.FormValue("description")
	difficultyLevelId := ctx.FormValue("diff_id")

	if name == "" {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      errors.New("name is required"),
			Message:    "name is required",
		}
	}

	if utf8.RuneCountInString(description) > 5000 {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      errors.New("maximum 5000 characters allowed"),
			Message:    "maximum 5000 characters allowed",
		}
	}

	privacy, err := strconv.ParseBool(ctx.FormValue("privacy"))
	if err != nil {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      err,
			Message:    err.Error(),
		}
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      err,
			Message:    err.Error(),
		}
	}

	files := form.File["images"]
	if len(files) > 10 {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      errors.New("maximum 10 images allowed"),
			Message:    "maximum 10 images allowed",
		}
	}

	quest := &model.Quest{
		Name:              name,
		Description:       description,
		DifficultyLevelID: difficultyLevelId,
		Private:           privacy,
	}

	if err := u.UserService.CreateQuest(ctx, quest, files); err != nil {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      err,
			Message:    err.Error(),
		}
	}

	return dto.HandlerResponse{
		StatusCode: fiber.StatusCreated,
		Message:    "Success",
	}
}

// CreateQuest godoc
// @Summary Create a new character
// @Description Allows a registered user to create a new character with images
// @Tags Character
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "Character Name"
// @Param description formData string false "Character Description (Max: 5000 digits)"
// @Param race_id formData string true "Race ID"
// @Param class_id formData string true "Class ID"
// @Param privacy formData boolean true "Character Privacy (true = private, false = public)"
// @Param images formData file false "Character Images (Max: 10 files)" collectionFormat=multi
// @Success 201 {object} dto.HandlerResponse
// @Failure 400 {object} dto.HandlerResponse
// @Router /user/character [post]
// @Security BearerAuth
func (u *UserHandler) CreateCharacter(ctx *fiber.Ctx) dto.HandlerResponse {
	name := ctx.FormValue("name")
	description := ctx.FormValue("description")
	raceID := ctx.FormValue("race_id")
	classID := ctx.FormValue("class_id")

	if name == "" {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      errors.New("name is required"),
			Message:    "name is required",
		}
	}

	if utf8.RuneCountInString(description) > 5000 {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      errors.New("maximum 5000 characters allowed"),
			Message:    "maximum 5000 characters allowed",
		}
	}

	privacy, err := strconv.ParseBool(ctx.FormValue("privacy"))
	if err != nil {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      err,
			Message:    err.Error(),
		}
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      err,
			Message:    err.Error(),
		}
	}

	files := form.File["images"]
	if len(files) > 10 {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      errors.New("maximum 10 images allowed"),
			Message:    "maximum 10 images allowed",
		}
	}

	character := &model.Character{
		Name:        name,
		Description: description,
		RaceID:      raceID,
		ClassID:     classID,
		Private:     privacy,
	}

	if err := u.UserService.CreateCharacter(ctx, character, files); err != nil {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      err,
			Message:    err.Error(),
		}
	}

	return dto.HandlerResponse{
		StatusCode: fiber.StatusCreated,
		Message:    "Success",
	}
}

// UpdateQuest godoc
// @Summary Update an existing quest
// @Description Updates a quest's details including name, description, difficulty level, privacy, and images
// @Tags Quest
// @Accept multipart/form-data
// @Produce json
// @Param quest_id path string false "Quest ID"
// @Param name formData string false "Quest Name"
// @Param description formData string false "Quest Description (max 5000 characters)"
// @Param diff_id formData string false "Difficulty Level ID"
// @Param privacy formData boolean false "Quest Privacy (true = private, false = public)"
// @Param images formData file false "Quest Images (Max: 10 files)" collectionFormat=multi
// @Success 200 {object} dto.HandlerResponse
// @Failure 400 {object} dto.HandlerResponse
// @Router /user/quest/{quest_id} [put]
// @Security BearerAuth
func (u *UserHandler) UpdateQuest(ctx *fiber.Ctx) dto.HandlerResponse {
	questID := ctx.Params("quest_id")

	if questID == "" {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      errors.New("quest_id is required"),
			Message:    "quest_id is required",
		}
	}

	name := ctx.FormValue("name")
	description := ctx.FormValue("description")
	difficultyLevelId := ctx.FormValue("diff_id")

	if utf8.RuneCountInString(description) > 5000 {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      errors.New("maximum 5000 characters allowed"),
			Message:    "maximum 5000 characters allowed",
		}
	}

	shouldUpdatePrivacy := true
	privacy, err := strconv.ParseBool(ctx.FormValue("privacy"))
	if err != nil {
		u.Logger.Info("Should not update privacy")
		shouldUpdatePrivacy = false
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      err,
			Message:    err.Error(),
		}
	}

	files := form.File["images"]
	if len(files) > 10 {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      errors.New("maximum 10 images allowed"),
			Message:    "maximum 10 images allowed",
		}
	}

	quest := &model.Quest{
		ID:                questID,
		Name:              name,
		Description:       description,
		DifficultyLevelID: difficultyLevelId,
		Private:           privacy,
	}

	if err := u.UserService.UpdateQuest(ctx, quest, files, shouldUpdatePrivacy); err != nil {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      err,
			Message:    err.Error(),
		}
	}

	return dto.HandlerResponse{
		StatusCode: fiber.StatusOK,
		Message:    "Success",
	}
}

// UpdateCharacter godoc
// @Summary Update an existing character
// @Description Updates a character's details including name, description, difficulty level, privacy, and images
// @Tags Character
// @Accept multipart/form-data
// @Produce json
// @Param character_id path string false "Character ID"
// @Param name formData string false "Character Name"
// @Param description formData string false "Character Description (max 5000 characters)"
// @Param race_id formData string false "Race ID"
// @Param class_id formData string false "Class ID"
// @Param privacy formData boolean false "Character Privacy (true = private, false = public)"
// @Param images formData file false "Character Images (Max: 10 files)" collectionFormat=multi
// @Success 200 {object} dto.HandlerResponse
// @Failure 400 {object} dto.HandlerResponse
// @Router /user/character/{character_id} [put]
// @Security BearerAuth
func (u *UserHandler) UpdateCharacter(ctx *fiber.Ctx) dto.HandlerResponse {
	characterID := ctx.Params("character_id")

	if characterID == "" {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      errors.New("character_id is required"),
			Message:    "character_id is required",
		}
	}

	name := ctx.FormValue("name")
	description := ctx.FormValue("description")
	raceID := ctx.FormValue("race_id")
	classID := ctx.FormValue("class_id")

	if utf8.RuneCountInString(description) > 5000 {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      errors.New("maximum 5000 characters allowed"),
			Message:    "maximum 5000 characters allowed",
		}
	}

	shouldUpdatePrivacy := true
	privacy, err := strconv.ParseBool(ctx.FormValue("privacy"))
	if err != nil {
		u.Logger.Info("Should not update privacy")
		shouldUpdatePrivacy = false
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      err,
			Message:    err.Error(),
		}
	}

	files := form.File["images"]
	if len(files) > 10 {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      errors.New("maximum 10 images allowed"),
			Message:    "maximum 10 images allowed",
		}
	}

	character := &model.Character{
		ID:          characterID,
		Name:        name,
		Description: description,
		RaceID:      raceID,
		ClassID:     classID,
		Private:     privacy,
	}

	if err := u.UserService.UpdateCharacter(ctx, character, files, shouldUpdatePrivacy); err != nil {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      err,
			Message:    err.Error(),
		}
	}

	return dto.HandlerResponse{
		StatusCode: fiber.StatusOK,
		Message:    "Success",
	}
}

// DeleteQuest godoc
// @Summary Delete a quest
// @Description Deletes a quest by its ID
// @Tags Quest
// @Param quest_id path string true "Quest ID"
// @Success 200 {object} dto.HandlerResponse "Success"
// @Failure 400 {object} dto.HandlerResponse "Bad Request"
// @Router /user/quest/{quest_id} [delete]
// @Security BearerAuth
func (u *UserHandler) DeleteQuest(ctx *fiber.Ctx) dto.HandlerResponse {
	questID := ctx.Params("quest_id")

	if questID == "" {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      errors.New("quest_id is required"),
			Message:    "quest_id is required",
		}
	}

	if err := u.UserService.DeleteQuest(ctx, questID); err != nil {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      err,
			Message:    err.Error(),
		}
	}

	return dto.HandlerResponse{
		StatusCode: fiber.StatusOK,
		Message:    "Success",
	}
}

// DeleteCharacter godoc
// @Summary Delete a character
// @Description Deletes a character by its ID
// @Tags Character
// @Param character_id path string true "Character ID"
// @Success 200 {object} dto.HandlerResponse "Success"
// @Failure 400 {object} dto.HandlerResponse "Bad Request"
// @Router /user/character/{character_id} [delete]
// @Security BearerAuth
func (u *UserHandler) DeleteCharacter(ctx *fiber.Ctx) dto.HandlerResponse {
	characterID := ctx.Params("character_id")

	if characterID == "" {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      errors.New("character_id is required"),
			Message:    "character_id is required",
		}
	}

	if err := u.UserService.DeleteQuest(ctx, characterID); err != nil {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      err,
			Message:    err.Error(),
		}
	}

	return dto.HandlerResponse{
		StatusCode: fiber.StatusOK,
		Message:    "Success",
	}
}
