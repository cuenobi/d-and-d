package handler

import (
	"log/slog"
	"net/url"
	"os"
	"path/filepath"

	"d-and-d/config"
	"d-and-d/internal/dto"
	"d-and-d/internal/port"

	"github.com/gofiber/fiber/v2"
)

type PublicHandler struct {
	Fiber         *fiber.App
	PublicService port.PublicService
	Logger        *slog.Logger
	Config        *config.Config
}

// RegisterHandler godoc
// @Summary Get public quests
// @Description Get public quests
// @Tags Public
// @Accept  json
// @Produce  json
// @Success 201 {object} dto.HandlerResponse
// @Failure 400 {object} dto.HandlerResponse
// @Router /public/quests [get]
func (p *PublicHandler) GetPublicQuest(ctx *fiber.Ctx) dto.HandlerResponse {
	quests, err := p.PublicService.GetPublicQuest()
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
// @Summary Get public characters
// @Description Get public characters
// @Tags Public
// @Accept  json
// @Produce  json
// @Success 201 {object} dto.HandlerResponse
// @Failure 400 {object} dto.HandlerResponse
// @Router /public/characters [get]
func (p *PublicHandler) GetPublicCharacter(ctx *fiber.Ctx) dto.HandlerResponse {
	characters, err := p.PublicService.GetPublicCharacter()
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

// GetImageFile godoc
// @Summary Get an image file
// @Description Returns an image file from the server
// @Tags Public
// @Param filename query string true "Image filename"
// @Produce octet-stream
// @Success 200 {file} binary "Image file"
// @Failure 400 {object} ErrorResponse "File not found"
// @Router /public/images [get]
func (p *PublicHandler) GetImageFile(ctx *fiber.Ctx) error {
	encodedFilename := ctx.Query("filename")

	filename, err := url.QueryUnescape(encodedFilename)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	filePath := filepath.Join("/app/images", filename)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.SendFile(filePath)
}
