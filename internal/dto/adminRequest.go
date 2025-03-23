package dto

type RaceBody struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required,max=5000"`
}

type ClassBody struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required,max=5000"`
}

type DifficultyLevel struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required,max=5000"`
}

type UpdateRaceBody struct {
	ID          string `json:"id" validate:"required"`
	Name        string `json:"name"`
	Description string `json:"description" validate:"max=5000"`
}

type UpdateClassBody struct {
	ID          string `json:"id" validate:"required"`
	Name        string `json:"name"`
	Description string `json:"description" validate:"max=5000"`
}

type UpdateDifficultyLevel struct {
	ID          string `json:"id" validate:"required"`
	Name        string `json:"name"`
	Description string `json:"description" validate:"max=5000"`
}
