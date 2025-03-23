package handler

import (
	"errors"
	"log/slog"

	"d-and-d/internal/dto"
	"d-and-d/internal/model"
	"d-and-d/internal/port"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AdminHandler struct {
	Fiber        *fiber.App
	JWT          port.JWT
	AdminService port.AdminService
	Logger       *slog.Logger
	Validator    *validator.Validate
}

// CreateRace godoc
// @Summary Create a new Race
// @Description Create a new Race
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param request body dto.RaceBody true "Create race details"
// @Success 201 {object} dto.HandlerResponse
// @Failure 400 {object} dto.HandlerResponse
// @Router /admin/race [post]
// @Security BearerAuth
func (a *AdminHandler) CreateRace(ctx *fiber.Ctx) dto.HandlerResponse {
	var input *dto.RaceBody

	if err := ctx.BodyParser(&input); err != nil {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      err,
			Message:    err.Error(),
		}
	}

	if err := a.Validator.Struct(input); err != nil {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      err,
			Message:    err.Error(),
		}
	}

	race := &model.Race{
		Name:        input.Name,
		Description: input.Description,
	}

	if err := a.AdminService.CreateRace(race); err != nil {
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

// CreateClass godoc
// @Summary Create a new Class
// @Description Create a new Class
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param request body dto.ClassBody true "Create class details"
// @Success 201 {object} dto.HandlerResponse
// @Failure 400 {object} dto.HandlerResponse
// @Router /admin/class [post]
// @Security BearerAuth
func (a *AdminHandler) CreateClass(ctx *fiber.Ctx) dto.HandlerResponse {
	var input *dto.ClassBody

	if err := ctx.BodyParser(&input); err != nil {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      err,
			Message:    err.Error(),
		}
	}

	if err := a.Validator.Struct(input); err != nil {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      err,
			Message:    err.Error(),
		}
	}

	class := &model.Class{
		Name:        input.Name,
		Description: input.Description,
	}

	if err := a.AdminService.CreateClass(class); err != nil {
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

// CreateDifficultyLevel godoc
// @Summary Create a new Difficulty Level
// @Description Create a new Difficulty Level
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param request body dto.DifficultyLevel true "Create difficultyLevel details"
// @Success 201 {object} dto.HandlerResponse
// @Failure 400 {object} dto.HandlerResponse
// @Router /admin/diff-lv [post]
// @Security BearerAuth
func (a *AdminHandler) CreateDifficultyLevel(ctx *fiber.Ctx) dto.HandlerResponse {
	var input *dto.DifficultyLevel

	if err := ctx.BodyParser(&input); err != nil {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      err,
			Message:    err.Error(),
		}
	}

	if err := a.Validator.Struct(input); err != nil {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      err,
			Message:    err.Error(),
		}
	}

	diff := &model.DifficultyLevels{
		Name:        input.Name,
		Description: input.Description,
	}

	if err := a.AdminService.CreateDifficultyLevel(diff); err != nil {
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

// UpdateRace godoc
// @Summary Update a new Race
// @Description Update a new Race
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param request body dto.UpdateRaceBody false "Update race details. If a field does not need to be updated, just insert an empty string, You can get your race_id from get race api"
// @Success 200 {object} dto.HandlerResponse
// @Failure 400 {object} dto.HandlerResponse
// @Router /admin/race [put]
// @Security BearerAuth
func (a *AdminHandler) UpdateRace(ctx *fiber.Ctx) dto.HandlerResponse {
	var input *dto.UpdateRaceBody

	if err := ctx.BodyParser(&input); err != nil {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      err,
			Message:    err.Error(),
		}
	}

	if err := a.Validator.Struct(input); err != nil {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      err,
			Message:    err.Error(),
		}
	}

	race := &model.Race{
		ID:          input.ID,
		Name:        input.Name,
		Description: input.Description,
	}

	if err := a.AdminService.UpdateRace(race); err != nil {
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

// UpdateClass godoc
// @Summary Update a new Class
// @Description Update a new Class
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param request body dto.UpdateClassBody false "Update class details. If a field does not need to be updated, just insert an empty string, You can get your class_id from get class api"
// @Success 200 {object} dto.HandlerResponse
// @Failure 400 {object} dto.HandlerResponse
// @Router /admin/class [put]
// @Security BearerAuth
func (a *AdminHandler) UpdateClass(ctx *fiber.Ctx) dto.HandlerResponse {
	var input *dto.UpdateClassBody

	if err := ctx.BodyParser(&input); err != nil {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      err,
			Message:    err.Error(),
		}
	}

	if err := a.Validator.Struct(input); err != nil {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      err,
			Message:    err.Error(),
		}
	}

	class := &model.Class{
		ID:          input.ID,
		Name:        input.Name,
		Description: input.Description,
	}

	if err := a.AdminService.UpdateClass(class); err != nil {
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

// UpdateDifficultyLevel godoc
// @Summary Update a new Difficulty Level
// @Description Update a new Difficulty Level
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param request body dto.UpdateDifficultyLevel false "Update Difficulty Levels. If a field does not need to be updated, just insert an empty string, You can get your difficulty_level_id from get difficulty_level api"
// @Success 200 {object} dto.HandlerResponse
// @Failure 400 {object} dto.HandlerResponse
// @Router /admin/diff-lv [put]
// @Security BearerAuth
func (a *AdminHandler) UpdateDifficultyLevel(ctx *fiber.Ctx) dto.HandlerResponse {
	var input *dto.UpdateDifficultyLevel

	if err := ctx.BodyParser(&input); err != nil {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      err,
			Message:    err.Error(),
		}
	}

	if err := a.Validator.Struct(input); err != nil {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      err,
			Message:    err.Error(),
		}
	}

	diff := &model.DifficultyLevels{
		ID:          input.ID,
		Name:        input.Name,
		Description: input.Description,
	}

	if err := a.AdminService.UpdateDifficultyLevel(diff); err != nil {
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

// DeleteRace godoc
// @Summary Delete a race
// @Description Deletes a race by its ID
// @Tags Admin
// @Param race_id path string true "race ID"
// @Success 200 {object} dto.HandlerResponse "Success"
// @Failure 400 {object} dto.HandlerResponse "Bad Request"
// @Router /admin/race/{race_id} [delete]
// @Security BearerAuth
func (a *AdminHandler) DeleteRace(ctx *fiber.Ctx) dto.HandlerResponse {
	raceID := ctx.Params("race_id")

	if raceID == "" {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      errors.New("race_id is required"),
			Message:    "race_id is required",
		}
	}

	if err := a.AdminService.DeleteRace(raceID); err != nil {
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

// DeleteClass godoc
// @Summary Delete a class
// @Description Deletes a class by its ID
// @Tags Admin
// @Param class_id path string true "class ID"
// @Success 200 {object} dto.HandlerResponse "Success"
// @Failure 400 {object} dto.HandlerResponse "Bad Request"
// @Router /admin/class/{class_id} [delete]
// @Security BearerAuth
func (a *AdminHandler) DeleteClass(ctx *fiber.Ctx) dto.HandlerResponse {
	raceID := ctx.Params("class_id")

	if raceID == "" {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      errors.New("class_id is required"),
			Message:    "class_id is required",
		}
	}

	if err := a.AdminService.DeleteClass(raceID); err != nil {
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

// DeleteDeleteDifficultyLevel godoc
// @Summary Delete a difficulty level
// @Description Deletes a difficulty level by its ID
// @Tags Admin
// @Param difficulty_level_id path string true "difficulty level ID"
// @Success 200 {object} dto.HandlerResponse "Success"
// @Failure 400 {object} dto.HandlerResponse "Bad Request"
// @Router /admin/diff-lv/{difficulty_level_id} [delete]
// @Security BearerAuth
func (a *AdminHandler) DeleteDifficultyLevel(ctx *fiber.Ctx) dto.HandlerResponse {
	raceID := ctx.Params("difficulty_level_id")

	if raceID == "" {
		return dto.HandlerResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      errors.New("difficulty_level_id is required"),
			Message:    "difficulty_level_id is required",
		}
	}

	if err := a.AdminService.DeleteDifficultyLevel(raceID); err != nil {
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
