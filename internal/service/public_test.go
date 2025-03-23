package service

import (
	"errors"
	"testing"

	"d-and-d/internal/dto"
	"d-and-d/internal/model"
	"d-and-d/mocks"

	"github.com/lib/pq"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestPublicServiceSuite(t *testing.T) {
	suite.Run(t, new(PublicServiceSuite))
}

type PublicServiceSuite struct {
	suite.Suite
	mockPublicRepo   *mocks.PublicRepository
	service          *PublicService
	mockCharacters   []*model.Character
	mockQuests       []*model.Quest
	expectCharacters []dto.PublicCharacterResponse
	expectQuests     []dto.PublicQuestResponse
}

func (p *PublicServiceSuite) SetupTest() {
	p.mockPublicRepo = mocks.NewPublicRepository(p.T())
	p.service = NewPublicService(p.mockPublicRepo)
	p.mockCharacters = []*model.Character{
		{
			Name:        "Aragorn",
			Description: "A ranger of the North",
			Class:       &model.Class{Name: "Ranger"},
			Race:        &model.Race{Name: "Human"},
			Images:      pq.StringArray{"image1.jpg", "image2.jpg"},
		},
		{
			Name:        "Legolas",
			Description: "An elven prince",
			Class:       &model.Class{Name: "Archer"},
			Race:        &model.Race{Name: "Elf"},
			Images:      pq.StringArray{"image3.jpg"},
		},
	}
	p.mockQuests = []*model.Quest{
		{
			Name:        "Defeat the Goblin King",
			Description: "A quest to defeat the Goblin King and bring peace to the village.",
			Private:     false,
			Images:      pq.StringArray{"image1.jpg", "image2.jpg"},
			DifficultyLevels: &model.DifficultyLevels{
				Name:        "Hard",
				Description: "Hell",
			},
		},
		{
			Name:        "Defeat the Dragon King",
			Description: "A quest to defeat the Dragon King and bring peace to the village.",
			Private:     true,
			Images:      pq.StringArray{"image1.jpg", "image2.jpg"},
			DifficultyLevels: &model.DifficultyLevels{
				Name:        "Easy",
				Description: "Chill",
			},
		},
	}
	p.expectCharacters = []dto.PublicCharacterResponse{
		{
			Name:        "Aragorn",
			Description: "A ranger of the North",
			ClassName:   "Ranger",
			RaceName:    "Human",
			Images:      pq.StringArray{"image1.jpg", "image2.jpg"},
		},
		{
			Name:        "Legolas",
			Description: "An elven prince",
			ClassName:   "Archer",
			RaceName:    "Elf",
			Images:      pq.StringArray{"image3.jpg"},
		},
	}
	p.expectQuests = []dto.PublicQuestResponse{
		{
			Name:            "Defeat the Goblin King",
			Description:     "A quest to defeat the Goblin King and bring peace to the village.",
			Difficulty:      "Hard",
			DiffDescription: "Hell",
			Images:          pq.StringArray{"image1.jpg", "image2.jpg"},
		},
		{
			Name:            "Defeat the Dragon King",
			Description:     "A quest to defeat the Dragon King and bring peace to the village.",
			Difficulty:      "Easy",
			DiffDescription: "Chill",
			Images:          pq.StringArray{"image1.jpg", "image2.jpg"},
		},
	}
}

func (p *PublicServiceSuite) TearDownTest() {
	p.mockPublicRepo.AssertExpectations(p.T())
}

func (p *PublicServiceSuite) SetupSubTest() {
	p.TearDownTest()
	p.SetupTest()
}

func (p *PublicServiceSuite) TestGetPublicCharacterSuccess() {
	p.mockPublicRepo.EXPECT().GetPublicCharacter().Return(p.mockCharacters, nil)
	result, err := p.service.GetPublicCharacter()
	p.Require().NoError(err)
	for i, character := range result {
		p.Equal(p.expectCharacters[i], character)
	}
}

func (p *PublicServiceSuite) TestGetPublicCharacterError() {
	p.mockPublicRepo.EXPECT().GetPublicCharacter().Return(nil, errors.New(mock.Anything))
	_, err := p.service.GetPublicCharacter()
	p.Require().Error(err)
}

func (p *PublicServiceSuite) TestGetPublicQuestSuccess() {
	p.mockPublicRepo.EXPECT().GetPublicQuest().Return(p.mockQuests, nil)
	result, err := p.service.GetPublicQuest()
	p.Require().NoError(err)
	for i, quest := range result {
		p.Equal(p.expectQuests[i], quest)
	}
}

func (p *PublicServiceSuite) TestGetPublicQuestError() {
	p.mockPublicRepo.EXPECT().GetPublicQuest().Return(nil, errors.New(mock.Anything))
	_, err := p.service.GetPublicQuest()
	p.Require().Error(err)
}
