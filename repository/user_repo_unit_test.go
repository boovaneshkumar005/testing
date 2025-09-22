package repository

import (
	"gotesting/config"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetUserByID_Unit(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Replace global DB with mock
	oldDB := config.DB
	config.DB = db
	defer func() { config.DB = oldDB }()

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "Alice")

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name FROM users WHERE id=?")).
		WithArgs(1).
		WillReturnRows(rows)

	user, err := GetUserByID(1)
	assert.NoError(t, err)
	assert.Equal(t, "Alice", user.Name)
}

func TestAddUser_Unit(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Replace global DB with mock
	oldDB := config.DB
	config.DB = db
	defer func() { config.DB = oldDB }()

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO users(id, name) VALUES(?, ?)")).
		WithArgs(2, "Bob").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = AddUser(User{ID: 2, Name: "Bob"})
	assert.NoError(t, err)
}
