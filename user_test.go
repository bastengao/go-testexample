package testexample

import (
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/khaiql/dbcleaner.v2"
	"gopkg.in/khaiql/dbcleaner.v2/engine"

	"github.com/bastengao/go-testexample/mock"
)

func TestRegisterUser(t *testing.T) {
	db := openDB()
	defer db.Close()

	option := dbcleaner.SetLockFileDir("./tmp/")
	cleaner := dbcleaner.New(option)
	sqlite := engine.NewSqliteEngine("test.db")
	cleaner.SetEngine(sqlite)

	cleaner.Acquire("users")
	defer cleaner.Clean("users")

	email := "gopher@golang.org"

	ctrl := gomock.NewController(t)
	mailer := mock.NewMockMailer(ctrl)
	mailer.
		EXPECT().
		Send(
			gomock.Eq(email),
			gomock.Any(),
			gomock.Any(),
		).
		Return(nil)

	err := registerUser(db, email, mailer)
	assert.NoError(t, err)
}

func TestCreateUser(t *testing.T) {
	db := openDB()
	defer db.Close()

	option := dbcleaner.SetLockFileDir("./tmp/")
	cleaner := dbcleaner.New(option)
	sqlite := engine.NewSqliteEngine("test.db")
	cleaner.SetEngine(sqlite)

	cleaner.Acquire("users")
	defer cleaner.Clean("users")

	err := createUser(db, "gopher@golang.org")
	assert.NoError(t, err)
}

func TestQueryUser(t *testing.T) {
	db := openDB()
	defer db.Close()

	fixtures, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("sqlite3"),
		testfixtures.Files("fixtures/users.yml"),
	)
	if err != nil {
		panic(err)
	}

	if err = fixtures.Load(); err != nil {
		panic(err)
	}

	user, err := queryUser(db, 1)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, int64(1), user.ID)
	assert.Equal(t, "gopher@golang.org", user.Email)
}
