package repository

import (
	"database/sql"
	"regexp"

	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SuiteSize struct {
	suite.Suite
	GormDB   *gorm.DB
	SqlDB    *sql.DB
	mock     sqlmock.Sqlmock
	sizeRepo SizeRepository
}

func SetupSuiteSize(t *testing.T) *SuiteSize {
	s := &SuiteSize{}
	var err error
	s.SqlDB, s.mock, err = sqlmock.New()
	if err != nil {
		t.Errorf("Failed to open mock sql db, got error: %v", err)
	}

	if s.SqlDB == nil {
		t.Errorf("mock db null")
	}

	if s.mock == nil {
		t.Error("sqlmock null")
	}

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 s.SqlDB,
		PreferSimpleProtocol: true,
	})
	s.GormDB, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Errorf("Failed to open gorm v2 db, got error: %v", err)
	}

	if s.GormDB == nil {
		t.Error("gorm db null")
	}
	s.sizeRepo = NewSizeRepository(s.GormDB)
	return s
}

func TestGetAllSizes(t *testing.T) {
	suite := SetupSuiteSize(t)

	query := `
	SELECT * FROM "sizes" 
	`

	user := sqlmock.NewRows([]string{"id", "name", "description", "price"}).
		AddRow(1, 2, 3, 4)

	suite.mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(1).WillReturnRows(user)

	data := suite.sizeRepo.AllSizes()

	// assert.NotNil()
	assert.Nil(t, data)

}
