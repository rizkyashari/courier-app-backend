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

type SuiteUser struct {
	suite.Suite
	GormDB   *gorm.DB
	SqlDB    *sql.DB
	mock     sqlmock.Sqlmock
	userRepo UserRepository
}

func SetupSuiteUser(t *testing.T) *SuiteUser {
	s := &SuiteUser{}
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
	s.userRepo = NewUserRepository(s.GormDB)
	return s
}

func TestGetAllUsers(t *testing.T) {
	suite := SetupSuiteUser(t)

	query := `
	SELECT * FROM "users"
	`

	user := sqlmock.NewRows([]string{"id", "user_id", "size_id", "address_id", "category_id", "payment_id", "add_on_id", "shipping_status", "review"}).
		AddRow(1, 2, 3, 4, 5, 6, 7, 8, 9)

	suite.mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(1).WillReturnRows(user)

	data := suite.userRepo.AllUsers()

	// assert.NotNil()
	assert.Nil(t, data)

}
