package testexample

import (
	"database/sql"
	"path/filepath"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/stretchr/testify/suite"
	"gopkg.in/khaiql/dbcleaner.v2"
	"gopkg.in/khaiql/dbcleaner.v2/engine"
)

func newCleaner() dbcleaner.DbCleaner {
	option := dbcleaner.SetLockFileDir("./tmp/")
	cleaner := dbcleaner.New(option)
	sqlite := engine.NewSqliteEngine("test.db")
	cleaner.SetEngine(sqlite)
	return cleaner
}

type baseSuite struct {
	suite.Suite
	cleaner      dbcleaner.DbCleaner
	fixtures     *testfixtures.Loader
	db           *sql.DB
	fixtureFiles []string
	tables       []string
}

func (s *baseSuite) SetupSuite() {
	s.cleaner = newCleaner()
	s.db = openDB()

	var newFixtureFiles []string
	for _, f := range s.fixtureFiles {
		newFixtureFiles = append(newFixtureFiles, filepath.Join("fixtures", f))
	}

	fixtures, err := testfixtures.New(
		testfixtures.Database(s.db),
		testfixtures.Dialect("sqlite3"),
		testfixtures.Files(newFixtureFiles...),
	)
	if err != nil {
		s.db.Close()
		panic(err)
	}
	s.fixtures = fixtures
}

func (s *baseSuite) SetupTest() {
	s.cleaner.Acquire(s.tables...)
	err := s.fixtures.Load()
	if err != nil {
		panic(err)
	}
}

func (s *baseSuite) TearDownTest() {
	s.cleaner.Clean(s.tables...)
}

func (s *baseSuite) TearDownSuite() {
	if s.cleaner != nil {
		s.cleaner.Close()
	}
	if s.db != nil {
		s.db.Close()
	}
}
