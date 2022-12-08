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

type SuiteAddress struct {
	suite.Suite
	GormDB      *gorm.DB
	SqlDB       *sql.DB
	mock        sqlmock.Sqlmock
	addressRepo AddressRepository
}

func SetupSuiteAddress(t *testing.T) *SuiteAddress {
	s := &SuiteAddress{}
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
	s.addressRepo = NewAddressRepository(s.GormDB)
	return s
}

func TestGetAllAddressesAdmin(t *testing.T) {
	suite := SetupSuiteAddress(t)

	query := `
	SELECT * FROM "addresses" WHERE "addresses"."deleted_at" IS NULL
	`

	user := sqlmock.NewRows([]string{"id", "user_id", "full_address", "recipient_name", "recipient_phone_number"}).
		AddRow(1, 2, 3, 4, 5)

	suite.mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(1).WillReturnRows(user)

	data := suite.addressRepo.AllAddressesAdmin()

	// assert.NotNil()
	assert.Nil(t, data)

}

func TestGetAllAddresses(t *testing.T) {
	suite := SetupSuiteAddress(t)

	var userID uint64

	query := `
	SELECT * FROM "addresses" WHERE "user_id" = ?
	`

	user := sqlmock.NewRows([]string{"id", "user_id", "full_address", "recipient_name", "recipient_phone_number"}).
		AddRow(1, 2, 3, 4, 5)

	suite.mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(1).WillReturnRows(user)

	data := suite.addressRepo.AllAddresses(userID)

	// assert.NotNil()
	assert.Nil(t, data)

}
