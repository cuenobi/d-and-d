package service

import (
	"bytes"
	"errors"
	"fmt"
	"log/slog"
	"net/http/httptest"
	"testing"

	"d-and-d/config"
	"d-and-d/internal/constant"
	"d-and-d/internal/dto"
	"d-and-d/internal/model"
	"d-and-d/mocks"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/valyala/fasthttp"
	"golang.org/x/crypto/bcrypt"
)

func TestUserServiceSuite(t *testing.T) {
	suite.Run(t, new(UserServiceSuite))
}

type UserServiceSuite struct {
	suite.Suite
	service                     *UserService
	mockUserRepo                mocks.UserRepository
	mockOptionRepo              mocks.OptionsRepository
	mockJwt                     mocks.JWT
	mockLogger                  *slog.Logger
	mockConfig                  *config.Config
	mockFiber                   *fiber.Ctx
	mockPassword                string
	mockUser                    *model.User
	mockAdmin                   *model.User
	mockRace                    []*model.Race
	mockClass                   []*model.Class
	mockDifficultyLevel         []*model.DifficultyLevels
	mockCharacters              []*model.Character
	mockQuests                  []*model.Quest
	mockExpectedRace            []dto.RaceResponse
	mockExpectedClass           []dto.ClassResponse
	mockExpectedDifficultyLevel []dto.DifficultyLevelResponse
	expectCharacters            []dto.CharacterResponse
	expectQuests                []dto.QuestResponse
}

func (u *UserServiceSuite) SetupTest() {
	u.mockUserRepo = *mocks.NewUserRepository(u.T())
	u.mockOptionRepo = *mocks.NewOptionsRepository(u.T())
	u.mockJwt = *mocks.NewJWT(u.T())

	u.mockPassword = "password"
	hashPwd, _ := bcrypt.GenerateFromPassword([]byte(u.mockPassword), bcrypt.DefaultCost)

	u.mockUser = &model.User{
		Username: "foo",
		Password: string(hashPwd),
		Name:     "bar",
		Role:     uint(constant.User),
	}

	u.mockAdmin = &model.User{
		Username: "bar",
		Password: string(hashPwd),
		Name:     "foo",
		Role:     uint(constant.Admin),
	}

	u.service = NewUserService(&u.mockUserRepo, &u.mockOptionRepo, &u.mockJwt, u.mockLogger, u.mockConfig)

	u.mockRace = []*model.Race{
		{
			ID:          "foo",
			Name:        "bar",
			Description: "foobar",
		},
	}
	u.mockClass = []*model.Class{
		{
			ID:          "foo",
			Name:        "bar",
			Description: "foobar",
		},
	}
	u.mockDifficultyLevel = []*model.DifficultyLevels{
		{
			ID:          "foo",
			Name:        "bar",
			Description: "foobar",
		},
	}

	u.mockExpectedRace = []dto.RaceResponse{
		{
			ID:          "foo",
			Name:        "bar",
			Description: "foobar",
		},
	}
	u.mockExpectedClass = []dto.ClassResponse{
		{
			ID:          "foo",
			Name:        "bar",
			Description: "foobar",
		},
	}
	u.mockExpectedDifficultyLevel = []dto.DifficultyLevelResponse{
		{
			ID:          "foo",
			Name:        "bar",
			Description: "foobar",
		},
	}
	u.mockCharacters = []*model.Character{
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
	u.mockQuests = []*model.Quest{
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
	u.expectCharacters = []dto.CharacterResponse{
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
	u.expectQuests = []dto.QuestResponse{
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

func (u *UserServiceSuite) TearDownTest() {
	u.mockUserRepo.AssertExpectations(u.T())
	u.mockJwt.AssertExpectations(u.T())
}

func (u *UserServiceSuite) SetupSubTest() {
	u.TearDownTest()
	u.SetupTest()
}

func (u *UserServiceSuite) TestCreateSuccess() {
	u.mockUserRepo.EXPECT().HasUsername(mock.Anything).Return(false, nil)
	u.mockUserRepo.EXPECT().CreateUser(mock.Anything).Return(nil)

	err := u.service.CreateUser(u.mockFiber, u.mockUser)
	u.Require().NoError(err)
}

func (u *UserServiceSuite) TestCreateInvalidRole() {
	err := u.service.CreateUser(u.mockFiber, u.mockAdmin)
	u.Require().Error(err)
	u.Require().Equal("invalid role", err.Error())
}

func (u *UserServiceSuite) TestCreatenameAlreadyExist() {
	u.mockUserRepo.EXPECT().HasUsername(mock.Anything).Return(true, nil)

	err := u.service.CreateUser(u.mockFiber, u.mockUser)
	u.Require().Error(err)
	u.Require().Equal("username already exist", err.Error())
}

func (u *UserServiceSuite) TestCreateHasUsernameError() {
	u.mockUserRepo.EXPECT().HasUsername(mock.Anything).Return(false, fmt.Errorf(mock.Anything))

	err := u.service.CreateUser(u.mockFiber, u.mockUser)
	u.Require().Error(err)
	u.Require().Equal(mock.Anything, err.Error())
}

func (u *UserServiceSuite) TestCreateCreateError() {
	u.mockUserRepo.EXPECT().HasUsername(mock.Anything).Return(false, nil)
	u.mockUserRepo.EXPECT().CreateUser(mock.Anything).Return(fmt.Errorf(mock.Anything))

	err := u.service.CreateUser(u.mockFiber, u.mockUser)
	u.Require().Error(err)
	u.Require().Equal(mock.Anything, err.Error())
}

func (u *UserServiceSuite) TestCreateAdminSuccess() {
	u.mockUserRepo.EXPECT().HasUsername(mock.Anything).Return(false, nil)
	u.mockUserRepo.EXPECT().CreateUser(mock.Anything).Return(nil)

	err := u.service.CreateAdmin(u.mockFiber, u.mockAdmin)
	u.Require().NoError(err)
}

func (u *UserServiceSuite) TestCreateAdminInvalidRole() {
	err := u.service.CreateAdmin(u.mockFiber, u.mockUser)
	u.Require().Error(err)
	u.Require().Equal("invalid role", err.Error())
}

func (u *UserServiceSuite) TestCreateAdminnameAlreadyExist() {
	u.mockUserRepo.EXPECT().HasUsername(mock.Anything).Return(true, nil)

	err := u.service.CreateAdmin(u.mockFiber, u.mockAdmin)
	u.Require().Error(err)
	u.Require().Equal("username already exist", err.Error())
}

func (u *UserServiceSuite) TestCreateAdminHasUsernameError() {
	u.mockUserRepo.EXPECT().HasUsername(mock.Anything).Return(false, fmt.Errorf(mock.Anything))

	err := u.service.CreateAdmin(u.mockFiber, u.mockAdmin)
	u.Require().Error(err)
	u.Require().Equal(mock.Anything, err.Error())
}

func (u *UserServiceSuite) TestCreateAdminCreateError() {
	u.mockUserRepo.EXPECT().HasUsername(mock.Anything).Return(false, nil)
	u.mockUserRepo.EXPECT().CreateUser(mock.Anything).Return(fmt.Errorf(mock.Anything))

	err := u.service.CreateAdmin(u.mockFiber, u.mockAdmin)
	u.Require().Error(err)
	u.Require().Equal(mock.Anything, err.Error())
}

func (u *UserServiceSuite) TestAuthenticationSuccess() {
	u.mockUserRepo.EXPECT().GetUserByUsername(mock.Anything).Return(u.mockUser, nil)
	u.mockJwt.EXPECT().Generate(mock.Anything, mock.Anything).Return(mock.Anything)

	_, _, err := u.service.Authentication(u.mockFiber, u.mockUser.Username, u.mockPassword)
	u.Require().NoError(err)
}

func (u *UserServiceSuite) TestAuthenticationGetError() {
	u.mockUserRepo.EXPECT().GetUserByUsername(mock.Anything).Return(nil, fmt.Errorf(mock.Anything))

	_, _, err := u.service.Authentication(u.mockFiber, u.mockUser.Username, u.mockPassword)
	u.Require().Error(err)
	u.Require().Equal(mock.Anything, err.Error())
}

func (u *UserServiceSuite) TestAuthenticationCompareError() {
	u.mockUserRepo.EXPECT().GetUserByUsername(mock.Anything).Return(u.mockUser, nil)

	_, _, err := u.service.Authentication(u.mockFiber, u.mockUser.Username, u.mockUser.Password)
	u.Require().Error(err)
}

func (u *UserServiceSuite) TestGetRaceSuccess() {
	u.mockOptionRepo.EXPECT().GetAllRace().Return(u.mockRace, nil)
	result, err := u.service.GetRace(u.mockFiber)
	u.Require().NoError(err)
	for i, r := range result {
		u.Equal(u.mockExpectedRace[i], r)
	}
}

func (u *UserServiceSuite) TestGetRaceError() {
	u.mockOptionRepo.EXPECT().GetAllRace().Return(nil, errors.New(mock.Anything))
	_, err := u.service.GetRace(u.mockFiber)
	u.Require().Error(err)
}

func (u *UserServiceSuite) TestGetClassSuccess() {
	u.mockOptionRepo.EXPECT().GetAllClass().Return(u.mockClass, nil)
	result, err := u.service.GetClass(u.mockFiber)
	u.Require().NoError(err)
	for i, r := range result {
		u.Equal(u.mockExpectedClass[i], r)
	}
}

func (u *UserServiceSuite) TestGetClassError() {
	u.mockOptionRepo.EXPECT().GetAllClass().Return(nil, errors.New(mock.Anything))
	_, err := u.service.GetClass(u.mockFiber)
	u.Require().Error(err)
}

func (u *UserServiceSuite) TestGetDifficultyLevelSuccess() {
	u.mockOptionRepo.EXPECT().GetAllDifficultyLevel().Return(u.mockDifficultyLevel, nil)
	result, err := u.service.GetDifficultyLevel(u.mockFiber)
	u.Require().NoError(err)
	for i, r := range result {
		u.Equal(u.mockExpectedDifficultyLevel[i], r)
	}
}

func (u *UserServiceSuite) TestGetDifficultyLevelError() {
	u.mockOptionRepo.EXPECT().GetAllDifficultyLevel().Return(nil, errors.New(mock.Anything))
	_, err := u.service.GetDifficultyLevel(u.mockFiber)
	u.Require().Error(err)
}

func (u *UserServiceSuite) TestGetCharactersSuccess() {
	u.mockUserRepo.EXPECT().GetAllCharacter().Return(u.mockCharacters, nil)
	result, err := u.service.GetAllCharacter(u.mockFiber)
	u.Require().NoError(err)
	for i, r := range result {
		u.Equal(u.expectCharacters[i], r)
	}
}

func (u *UserServiceSuite) TestGetCharactersError() {
	u.mockUserRepo.EXPECT().GetAllCharacter().Return(nil, errors.New(mock.Anything))
	_, err := u.service.GetAllCharacter(u.mockFiber)
	u.Require().Error(err)
}

func (u *UserServiceSuite) TestGetQuestsSuccess() {
	u.mockUserRepo.EXPECT().GetAllQuest().Return(u.mockQuests, nil)
	result, err := u.service.GetAllQuest(u.mockFiber)
	u.Require().NoError(err)
	for i, r := range result {
		u.Equal(u.expectQuests[i], r)
	}
}

func (u *UserServiceSuite) TestGetQuestsError() {
	u.mockUserRepo.EXPECT().GetAllQuest().Return(nil, errors.New(mock.Anything))
	_, err := u.service.GetAllQuest(u.mockFiber)
	u.Require().Error(err)
}

func (u *UserServiceSuite) TestCreateCharacterSuccess() {
	app := fiber.New()

	req := httptest.NewRequest("POST", "/character", bytes.NewBuffer([]byte(`{}`)))
	req.Header.Set("Content-Type", "application/json")

	reqCtx := &fasthttp.RequestCtx{}
	reqCtx.Request.SetRequestURI(req.URL.String())
	reqCtx.Request.Header.SetContentType(req.Header.Get("Content-Type"))
	reqCtx.Request.Header.SetMethod(req.Method)
	reqCtx.Request.SetBodyStream(req.Body, -1)

	ctx := app.AcquireCtx(reqCtx)
	defer app.ReleaseCtx(ctx)

	ctx.Locals("username", u.mockUser.Username)
	ctx.Locals("role", u.mockUser.Role)

	u.mockOptionRepo.EXPECT().GetClassByID(mock.Anything).Return(u.mockClass[0], nil)
	u.mockOptionRepo.EXPECT().GetRaceByID(mock.Anything).Return(u.mockRace[0], nil)
	u.mockUserRepo.EXPECT().GetUserByUsername(mock.Anything).Return(u.mockUser, nil)
	u.mockUserRepo.EXPECT().CreateCharacter(mock.Anything).Return(nil)

	err := u.service.CreateCharacter(ctx, u.mockCharacters[0], nil)
	u.Require().NoError(err)
}

func (u *UserServiceSuite) TestCreateQuestSuccess() {
	app := fiber.New()

	req := httptest.NewRequest("POST", "/character", bytes.NewBuffer([]byte(`{}`)))
	req.Header.Set("Content-Type", "application/json")

	reqCtx := &fasthttp.RequestCtx{}
	reqCtx.Request.SetRequestURI(req.URL.String())
	reqCtx.Request.Header.SetContentType(req.Header.Get("Content-Type"))
	reqCtx.Request.Header.SetMethod(req.Method)
	reqCtx.Request.SetBodyStream(req.Body, -1)

	ctx := app.AcquireCtx(reqCtx)
	defer app.ReleaseCtx(ctx)

	ctx.Locals("username", u.mockUser.Username)
	ctx.Locals("role", u.mockUser.Role)

	u.mockOptionRepo.EXPECT().GetDifficultyLevelById(mock.Anything).Return(u.mockDifficultyLevel[0], nil)
	u.mockUserRepo.EXPECT().GetUserByUsername(mock.Anything).Return(u.mockUser, nil)
	u.mockUserRepo.EXPECT().CreateQuest(mock.Anything).Return(nil)

	err := u.service.CreateQuest(ctx, u.mockQuests[0], nil)
	u.Require().NoError(err)
}

func (u *UserServiceSuite) TestUpdateCharacterSuccess() {
	app := fiber.New()

	req := httptest.NewRequest("POST", "/character", bytes.NewBuffer([]byte(`{}`)))
	req.Header.Set("Content-Type", "application/json")

	reqCtx := &fasthttp.RequestCtx{}
	reqCtx.Request.SetRequestURI(req.URL.String())
	reqCtx.Request.Header.SetContentType(req.Header.Get("Content-Type"))
	reqCtx.Request.Header.SetMethod(req.Method)
	reqCtx.Request.SetBodyStream(req.Body, -1)

	ctx := app.AcquireCtx(reqCtx)
	defer app.ReleaseCtx(ctx)

	ctx.Locals("username", u.mockUser.Username)
	ctx.Locals("role", u.mockUser.Role)

	u.mockOptionRepo.EXPECT().GetClassByID(mock.Anything).Return(u.mockClass[0], nil)
	u.mockOptionRepo.EXPECT().GetRaceByID(mock.Anything).Return(u.mockRace[0], nil)
	u.mockUserRepo.EXPECT().GetUserByUsername(mock.Anything).Return(u.mockUser, nil)
	u.mockUserRepo.EXPECT().UpdateCharacter(mock.Anything, mock.Anything, mock.Anything).Return(nil)

	err := u.service.UpdateCharacter(ctx, u.mockCharacters[0], nil, false)
	u.Require().NoError(err)
}

func (u *UserServiceSuite) TestUpdateQuestSuccess() {
	app := fiber.New()

	req := httptest.NewRequest("POST", "/character", bytes.NewBuffer([]byte(`{}`)))
	req.Header.Set("Content-Type", "application/json")

	reqCtx := &fasthttp.RequestCtx{}
	reqCtx.Request.SetRequestURI(req.URL.String())
	reqCtx.Request.Header.SetContentType(req.Header.Get("Content-Type"))
	reqCtx.Request.Header.SetMethod(req.Method)
	reqCtx.Request.SetBodyStream(req.Body, -1)

	ctx := app.AcquireCtx(reqCtx)
	defer app.ReleaseCtx(ctx)

	ctx.Locals("username", u.mockUser.Username)
	ctx.Locals("role", u.mockUser.Role)

	u.mockOptionRepo.EXPECT().GetDifficultyLevelById(mock.Anything).Return(u.mockDifficultyLevel[0], nil)
	u.mockUserRepo.EXPECT().GetUserByUsername(mock.Anything).Return(u.mockUser, nil)
	u.mockUserRepo.EXPECT().UpdateQuest(mock.Anything, mock.Anything, mock.Anything).Return(nil)

	err := u.service.UpdateQuest(ctx, u.mockQuests[0], nil, false)
	u.Require().NoError(err)
}

func (u *UserServiceSuite) TestDeleteCharacter() {
	app := fiber.New()

	req := httptest.NewRequest("POST", "/character", bytes.NewBuffer([]byte(`{}`)))
	req.Header.Set("Content-Type", "application/json")

	reqCtx := &fasthttp.RequestCtx{}
	reqCtx.Request.SetRequestURI(req.URL.String())
	reqCtx.Request.Header.SetContentType(req.Header.Get("Content-Type"))
	reqCtx.Request.Header.SetMethod(req.Method)
	reqCtx.Request.SetBodyStream(req.Body, -1)

	ctx := app.AcquireCtx(reqCtx)
	defer app.ReleaseCtx(ctx)

	ctx.Locals("username", u.mockUser.Username)
	ctx.Locals("role", u.mockUser.Role)

	u.mockUserRepo.EXPECT().GetUserByUsername(mock.Anything).Return(u.mockUser, nil)
	u.mockUserRepo.EXPECT().DeleteCharacter(mock.Anything, mock.Anything).Return(nil)

	err := u.service.DeleteCharacter(ctx, mock.Anything)
	u.Require().NoError(err)
}

func (u *UserServiceSuite) TestDeleteQuest() {
	app := fiber.New()

	req := httptest.NewRequest("POST", "/character", bytes.NewBuffer([]byte(`{}`)))
	req.Header.Set("Content-Type", "application/json")

	reqCtx := &fasthttp.RequestCtx{}
	reqCtx.Request.SetRequestURI(req.URL.String())
	reqCtx.Request.Header.SetContentType(req.Header.Get("Content-Type"))
	reqCtx.Request.Header.SetMethod(req.Method)
	reqCtx.Request.SetBodyStream(req.Body, -1)

	ctx := app.AcquireCtx(reqCtx)
	defer app.ReleaseCtx(ctx)

	ctx.Locals("username", u.mockUser.Username)
	ctx.Locals("role", u.mockUser.Role)

	u.mockUserRepo.EXPECT().GetUserByUsername(mock.Anything).Return(u.mockUser, nil)
	u.mockUserRepo.EXPECT().DeleteQuest(mock.Anything, mock.Anything).Return(nil)

	err := u.service.DeleteQuest(ctx, mock.Anything)
	u.Require().NoError(err)
}
