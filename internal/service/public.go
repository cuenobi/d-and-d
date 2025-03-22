package service

import (
	"d-and-d/internal/dto"
	"d-and-d/internal/port"
)

type PublicService struct {
	publicRepo port.PublicRepository
}

func NewPublicService(publicRepo port.PublicRepository) *PublicService {
	return &PublicService{
		publicRepo: publicRepo,
	}
}

func (p *PublicService) GetPublicCharacter() ([]dto.PublicCharacterResponse, error) {
	resp := make([]dto.PublicCharacterResponse, 0)
	characters, err := p.publicRepo.GetPublicCharacter()
	if err != nil {
		return nil, err
	}
	for _, character := range characters {
		resp = append(resp, dto.PublicCharacterResponse{
			Name:        character.Name,
			Description: character.Description,
			ClassName:   character.Class.Name,
			RaceName:    character.Race.Name,
			Images:      character.Images,
		})
	}
	return resp, nil
}

func (p *PublicService) GetPublicQuest() ([]dto.PublicQuestResponse, error) {
	resp := make([]dto.PublicQuestResponse, 0)
	quests, err := p.publicRepo.GetPublicQuest()
	if err != nil {
		return nil, err
	}
	for _, quest := range quests {
		resp = append(resp, dto.PublicQuestResponse{
			Name:            quest.Name,
			Description:     quest.Description,
			Difficulty:      quest.DifficultyLevels.Name,
			DiffDescription: quest.DifficultyLevels.Description,
			Images:          quest.Images,
		})
	}
	return resp, nil
}
