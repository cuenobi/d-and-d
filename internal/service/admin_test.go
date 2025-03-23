package service

import (
	"errors"
	"log/slog"
	"testing"

	"d-and-d/internal/model"
	"d-and-d/mocks"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestAdminServiceSuite(t *testing.T) {
	suite.Run(t, new(AdminServiceSuite))
}

type AdminServiceSuite struct {
	suite.Suite
	mockOptionRepo      *mocks.OptionsRepository
	service             *AdminService
	mockLogger          *slog.Logger
	mockRace            *model.Race
	mockClass           *model.Class
	mockDifficultyLevel *model.DifficultyLevels
}

func (a *AdminServiceSuite) SetupTest() {
	a.mockOptionRepo = mocks.NewOptionsRepository(a.T())
	a.service = NewAdminService(a.mockLogger, a.mockOptionRepo)
	a.mockRace = &model.Race{}
	a.mockClass = &model.Class{}
	a.mockDifficultyLevel = &model.DifficultyLevels{}
}

func (a *AdminServiceSuite) TearDownTest() {
	a.mockOptionRepo.AssertExpectations(a.T())
}

func (a *AdminServiceSuite) SetupSubTest() {
	a.TearDownTest()
	a.SetupTest()
}

func (a *AdminServiceSuite) TestCreateRaceSuccess() {
	a.mockOptionRepo.EXPECT().CreateRace(mock.Anything).Return(nil)
	err := a.service.CreateRace(a.mockRace)
	a.Require().NoError(err)
}

func (a *AdminServiceSuite) TestCreateRaceError() {
	a.mockOptionRepo.EXPECT().CreateRace(mock.Anything).Return(errors.New(mock.Anything))
	err := a.service.CreateRace(a.mockRace)
	a.Require().Error(err)
}

func (a *AdminServiceSuite) TestCreateClassSuccess() {
	a.mockOptionRepo.EXPECT().CreateClass(mock.Anything).Return(nil)
	err := a.service.CreateClass(a.mockClass)
	a.Require().NoError(err)
}

func (a *AdminServiceSuite) TestCreateClassError() {
	a.mockOptionRepo.EXPECT().CreateClass(mock.Anything).Return(errors.New(mock.Anything))
	err := a.service.CreateClass(a.mockClass)
	a.Require().Error(err)
}

func (a *AdminServiceSuite) TestCreateDifficultyLevelSuccess() {
	a.mockOptionRepo.EXPECT().CreateDifficultyLevel(mock.Anything).Return(nil)
	err := a.service.CreateDifficultyLevel(a.mockDifficultyLevel)
	a.Require().NoError(err)
}

func (a *AdminServiceSuite) TestCreateDifficultyLevelError() {
	a.mockOptionRepo.EXPECT().CreateDifficultyLevel(mock.Anything).Return(errors.New(mock.Anything))
	err := a.service.CreateDifficultyLevel(a.mockDifficultyLevel)
	a.Require().Error(err)
}

func (a *AdminServiceSuite) TestUpdateRaceSuccess() {
	a.mockOptionRepo.EXPECT().UpdateRace(mock.Anything).Return(nil)
	err := a.service.UpdateRace(a.mockRace)
	a.Require().NoError(err)
}

func (a *AdminServiceSuite) TestUpdateRaceError() {
	a.mockOptionRepo.EXPECT().UpdateRace(mock.Anything).Return(errors.New(mock.Anything))
	err := a.service.UpdateRace(a.mockRace)
	a.Require().Error(err)
}

func (a *AdminServiceSuite) TestUpdateClassSuccess() {
	a.mockOptionRepo.EXPECT().UpdateClass(mock.Anything).Return(nil)
	err := a.service.UpdateClass(a.mockClass)
	a.Require().NoError(err)
}

func (a *AdminServiceSuite) TestUpdateClassError() {
	a.mockOptionRepo.EXPECT().UpdateClass(mock.Anything).Return(errors.New(mock.Anything))
	err := a.service.UpdateClass(a.mockClass)
	a.Require().Error(err)
}

func (a *AdminServiceSuite) TestUpdateDifficultyLevelSuccess() {
	a.mockOptionRepo.EXPECT().UpdateDifficultyLevel(mock.Anything).Return(nil)
	err := a.service.UpdateDifficultyLevel(a.mockDifficultyLevel)
	a.Require().NoError(err)
}

func (a *AdminServiceSuite) TestUpdateDifficultyLevelError() {
	a.mockOptionRepo.EXPECT().UpdateDifficultyLevel(mock.Anything).Return(errors.New(mock.Anything))
	err := a.service.UpdateDifficultyLevel(a.mockDifficultyLevel)
	a.Require().Error(err)
}

func (a *AdminServiceSuite) TestDeleteRaceSuccess() {
	a.mockOptionRepo.EXPECT().DeleteRace(mock.Anything).Return(nil)
	err := a.service.DeleteRace(mock.Anything)
	a.Require().NoError(err)
}

func (a *AdminServiceSuite) TestDeleteRaceError() {
	a.mockOptionRepo.EXPECT().DeleteRace(mock.Anything).Return(errors.New(mock.Anything))
	err := a.service.DeleteRace(mock.Anything)
	a.Require().Error(err)
}

func (a *AdminServiceSuite) TestDeleteClassSuccess() {
	a.mockOptionRepo.EXPECT().DeleteClass(mock.Anything).Return(nil)
	err := a.service.DeleteClass(mock.Anything)
	a.Require().NoError(err)
}

func (a *AdminServiceSuite) TestDeleteClassError() {
	a.mockOptionRepo.EXPECT().DeleteClass(mock.Anything).Return(errors.New(mock.Anything))
	err := a.service.DeleteClass(mock.Anything)
	a.Require().Error(err)
}

func (a *AdminServiceSuite) TestDeleteDifficultyLevelSuccess() {
	a.mockOptionRepo.EXPECT().DeleteDifficultyLevel(mock.Anything).Return(nil)
	err := a.service.DeleteDifficultyLevel(mock.Anything)
	a.Require().NoError(err)
}

func (a *AdminServiceSuite) TestDeleteDifficultyLevelError() {
	a.mockOptionRepo.EXPECT().DeleteDifficultyLevel(mock.Anything).Return(errors.New(mock.Anything))
	err := a.service.DeleteDifficultyLevel(mock.Anything)
	a.Require().Error(err)
}
