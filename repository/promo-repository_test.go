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

type SuitePromo struct {
	suite.Suite
	GormDB    *gorm.DB
	SqlDB     *sql.DB
	mock      sqlmock.Sqlmock
	promoRepo PromoRepository
}

func SetupSuitePromo(t *testing.T) *SuitePromo {
	s := &SuitePromo{}
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
	s.promoRepo = NewPromoRepository(s.GormDB)
	return s
}

func TestGetAllPromos(t *testing.T) {
	suite := SetupSuitePromo(t)

	query := `
	SELECT * FROM "promos" 
	`

	user := sqlmock.NewRows([]string{"id", "name", "min_fee", "max_discount", "discount", "quota"}).
		AddRow(1, 2, 3, 4, 5, 6)

	suite.mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(1).WillReturnRows(user)

	data := suite.promoRepo.AllPromos()

	// assert.NotNil()
	assert.Nil(t, data)

}
