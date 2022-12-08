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

type SuiteUserPromo struct {
	suite.Suite
	GormDB        *gorm.DB
	SqlDB         *sql.DB
	mock          sqlmock.Sqlmock
	userPromoRepo UserPromoRepository
}

func SetupSuiteUserPromo(t *testing.T) *SuiteUserPromo {
	s := &SuiteUserPromo{}
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
	s.userPromoRepo = NewUserPromoRepository(s.GormDB)
	return s
}

func TestGetAllUserPromosAdmin(t *testing.T) {
	suite := SetupSuiteUserPromo(t)

	query := `
	SELECT * FROM "user_promos"
	`

	user := sqlmock.NewRows([]string{"id", "user_id", "promo_id", "status"}).
		AddRow(1, 2, 3, 4)

	suite.mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(1).WillReturnRows(user)

	data := suite.userPromoRepo.AllUserPromosAdmin()

	// assert.NotNil()
	assert.Nil(t, data)

}

func TestGetAllUserPromos(t *testing.T) {
	suite := SetupSuiteUserPromo(t)

	var userID uint64

	query := `
	SELECT * FROM "user_promos" WHERE "user_id" = ?
	`

	user := sqlmock.NewRows([]string{"id", "user_id", "promo_id", "status"}).
		AddRow(1, 2, 3, 4)

	suite.mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(1).WillReturnRows(user)

	data := suite.userPromoRepo.AllUserPromos(userID)

	// assert.NotNil()
	assert.Nil(t, data)

}
