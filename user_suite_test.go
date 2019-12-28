package testexample

import (
	"testing"

	"github.com/bastengao/go-testexample/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type UserTestSuite struct {
	baseSuite
}

func (s *UserTestSuite) TestRegsiterUser() {
	email := "test@golang.org"

	ctrl := gomock.NewController(s.T())
	mailer := mock.NewMockMailer(ctrl)
	mailer.
		EXPECT().
		Send(
			gomock.Eq(email),
			gomock.Any(),
			gomock.Any(),
		).
		Return(nil)

	err := registerUser(s.db, email, mailer)
	s.Assert().NoError(err)
}

func (s *UserTestSuite) TestCreateUser() {
	email := "test@golang.org"
	err := createUser(s.db, email)
	s.Assert().NoError(err)
}

func (s *UserTestSuite) TestQueryUser() {
	user, err := queryUser(s.db, 1)
	s.Assert().NoError(err)
	s.Assert().NotNil(user)
	s.Assert().Equal(int64(1), user.ID)
	s.Assert().Equal("gopher@golang.org", user.Email)
}

func TestUserTestSuite(t *testing.T) {
	s := &UserTestSuite{
		baseSuite: baseSuite{
			fixtureFiles: []string{"users.yml"},
			tables:       []string{"users"},
		},
	}

	suite.Run(t, s)
}
