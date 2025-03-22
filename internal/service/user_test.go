package service

// import (
// 	"fmt"
// 	"testing"

// 	"d-and-d/internal/constant"
// 	"d-and-d/internal/model"
// 	"d-and-d/mocks"

// 	"github.com/stretchr/testify/mock"
// 	"github.com/stretchr/testify/suite"
// 	"golang.org/x/crypto/bcrypt"
// )

// func TestServiceSuite(t *testing.T) {
// 	suite.Run(t, new(ServiceSuite))
// }

// type ServiceSuite struct {
// 	suite.Suite
// 	service      *UserService
// 	mockRepo     mocks.UserRepository
// 	mockJwt      mocks.JWT
// 	mock         *model.User
// 	mockAdmin    *model.User
// 	mockPassword string
// }

// func (u *ServiceSuite) SetupTest() {
// 	u.mockRepo = *mocks.NewUserRepository(u.T())
// 	u.mockJwt = *mocks.NewJWT(u.T())

// 	u.mockPassword = "password"
// 	hashPwd, _ := bcrypt.GenerateFromPassword([]byte(u.mockPassword), bcrypt.DefaultCost)

// 	u.mock = &model.User{
// 		Username: "foo",
// 		Password: string(hashPwd),
// 		Name:     "bar",
// 		Role:     uint(constant.User),
// 	}

// 	u.mockAdmin = &model.User{
// 		Username: "bar",
// 		Password: string(hashPwd),
// 		Name:     "foo",
// 		Role:     uint(constant.Admin),
// 	}

// 	u.service = NewUserService(&u.mockRepo, &u.mockJwt)
// }

// func (u *ServiceSuite) TearDownTest() {
// 	u.mockRepo.AssertExpectations(u.T())
// 	u.mockJwt.AssertExpectations(u.T())
// }

// func (u *ServiceSuite) SetupSubTest() {
// 	u.TearDownTest()
// 	u.SetupTest()
// }

// func (u *ServiceSuite) TestCreateSuccess() {
// 	u.mockRepo.EXPECT().HasUsername(mock.Anything).Return(false, nil)
// 	u.mockRepo.EXPECT().CreateUser(mock.Anything).Return(nil)

// 	err := u.service.CreateUser(u.mock)
// 	u.Require().NoError(err)
// }

// func (u *ServiceSuite) TestCreateInvalidRole() {
// 	err := u.service.CreateUser(u.mockAdmin)
// 	u.Require().Error(err)
// 	u.Require().Equal("invalid role", err.Error())
// }

// func (u *ServiceSuite) TestCreatenameAlreadyExist() {
// 	u.mockRepo.EXPECT().HasUsername(mock.Anything).Return(true, nil)

// 	err := u.service.CreateUser(u.mock)
// 	u.Require().Error(err)
// 	u.Require().Equal("username already exist", err.Error())
// }

// func (u *ServiceSuite) TestCreateHasUsernameError() {
// 	u.mockRepo.EXPECT().HasUsername(mock.Anything).Return(false, fmt.Errorf(mock.Anything))

// 	err := u.service.CreateUser(u.mock)
// 	u.Require().Error(err)
// 	u.Require().Equal(mock.Anything, err.Error())
// }

// func (u *ServiceSuite) TestCreateCreateError() {
// 	u.mockRepo.EXPECT().HasUsername(mock.Anything).Return(false, nil)
// 	u.mockRepo.EXPECT().CreateUser(mock.Anything).Return(fmt.Errorf(mock.Anything))

// 	err := u.service.CreateUser(u.mock)
// 	u.Require().Error(err)
// 	u.Require().Equal(mock.Anything, err.Error())
// }

// func (u *ServiceSuite) TestCreateAdminSuccess() {
// 	u.mockRepo.EXPECT().HasUsername(mock.Anything).Return(false, nil)
// 	u.mockRepo.EXPECT().CreateUser(mock.Anything).Return(nil)

// 	err := u.service.CreateAdmin(u.mockAdmin)
// 	u.Require().NoError(err)
// }

// func (u *ServiceSuite) TestCreateAdminInvalidRole() {
// 	err := u.service.CreateAdmin(u.mock)
// 	u.Require().Error(err)
// 	u.Require().Equal("invalid role", err.Error())
// }

// func (u *ServiceSuite) TestCreateAdminnameAlreadyExist() {
// 	u.mockRepo.EXPECT().HasUsername(mock.Anything).Return(true, nil)

// 	err := u.service.CreateAdmin(u.mockAdmin)
// 	u.Require().Error(err)
// 	u.Require().Equal("username already exist", err.Error())
// }

// func (u *ServiceSuite) TestCreateAdminHasUsernameError() {
// 	u.mockRepo.EXPECT().HasUsername(mock.Anything).Return(false, fmt.Errorf(mock.Anything))

// 	err := u.service.CreateAdmin(u.mockAdmin)
// 	u.Require().Error(err)
// 	u.Require().Equal(mock.Anything, err.Error())
// }

// func (u *ServiceSuite) TestCreateAdminCreateError() {
// 	u.mockRepo.EXPECT().HasUsername(mock.Anything).Return(false, nil)
// 	u.mockRepo.EXPECT().CreateUser(mock.Anything).Return(fmt.Errorf(mock.Anything))

// 	err := u.service.CreateAdmin(u.mockAdmin)
// 	u.Require().Error(err)
// 	u.Require().Equal(mock.Anything, err.Error())
// }

// func (u *ServiceSuite) TestAuthenticationSuccess() {
// 	u.mockRepo.EXPECT().GetUserByUsername(mock.Anything).Return(u.mock, nil)
// 	u.mockJwt.EXPECT().Generate(mock.Anything, mock.Anything).Return(mock.Anything)

// 	_, _, err := u.service.Authentication(u.mock.Username, u.mockPassword)
// 	u.Require().NoError(err)
// }

// func (u *ServiceSuite) TestAuthenticationGetError() {
// 	u.mockRepo.EXPECT().GetUserByUsername(mock.Anything).Return(nil, fmt.Errorf(mock.Anything))

// 	_, _, err := u.service.Authentication(u.mock.Username, u.mockPassword)
// 	u.Require().Error(err)
// 	u.Require().Equal(mock.Anything, err.Error())
// }

// func (u *ServiceSuite) TestAuthenticationCompareError() {
// 	u.mockRepo.EXPECT().GetUserByUsername(mock.Anything).Return(u.mock, nil)

// 	_, _, err := u.service.Authentication(u.mock.Username, u.mock.Password)
// 	u.Require().Error(err)
// }
