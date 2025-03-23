package service

import (
	"log/slog"

	"d-and-d/internal/model"
	"d-and-d/internal/port"
)

type AdminService struct {
	logger     *slog.Logger
	optionRepo port.OptionsRepository
}

func NewAdminService(logger *slog.Logger, optionRepo port.OptionsRepository) *AdminService {
	return &AdminService{
		logger:     logger,
		optionRepo: optionRepo,
	}
}

func (a *AdminService) CreateClass(class *model.Class) error {
	return a.optionRepo.CreateClass(class)
}

func (a *AdminService) CreateRace(race *model.Race) error {
	return a.optionRepo.CreateRace(race)
}

func (a *AdminService) CreateDifficultyLevel(diffLv *model.DifficultyLevels) error {
	return a.optionRepo.CreateDifficultyLevel(diffLv)
}

func (a *AdminService) UpdateRace(race *model.Race) error {
	return a.optionRepo.UpdateRace(race)
}

func (a *AdminService) UpdateClass(class *model.Class) error {
	return a.optionRepo.UpdateClass(class)
}

func (a *AdminService) UpdateDifficultyLevel(diff *model.DifficultyLevels) error {
	return a.optionRepo.UpdateDifficultyLevel(diff)
}

func (a *AdminService) DeleteRace(id string) error {
	return a.optionRepo.DeleteRace(id)
}

func (a *AdminService) DeleteClass(id string) error {
	return a.optionRepo.DeleteClass(id)
}

func (a *AdminService) DeleteDifficultyLevel(id string) error {
	return a.optionRepo.DeleteDifficultyLevel(id)
}
