package repository

import (
	"regexp"
	"testing"

	"github.com/minand-mohan/library-app-api/database/models"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestCreateUser(t *testing.T) {
	//initialize variables
	test_email := "test@example.com"
	test_username := "test"
	test_phone := "1234567890"
	test_id := "123e4567-e89b-12d3-a456-426614174000"

	tc := []struct {
		name          string
		user          *models.User
		mockFunction  func(mock sqlmock.Sqlmock, user *models.User) error
		expectedError error
	}{
		{
			name: "User created successfully",
			user: &models.User{
				Email:    &test_email,
				Phone:    &test_phone,
				Username: &test_username,
			},
			mockFunction: func(mock sqlmock.Sqlmock, user *models.User) error {
				query := regexp.QuoteMeta(`INSERT INTO "users" ("username","email","phone") VALUES ($1,$2,$3) RETURNING "id"`)
				mock.ExpectBegin()
				mock.ExpectQuery(query).
					WithArgs(*user.Username, *user.Email, *user.Phone).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(test_id))
				mock.ExpectCommit()
				return nil
			},
			expectedError: nil,
		},
		{
			name: "User creation failed",
			user: &models.User{
				Email:    &test_email,
				Phone:    &test_phone,
				Username: &test_username,
			},
			mockFunction: func(mock sqlmock.Sqlmock, user *models.User) error {
				err := sqlmock.ErrCancelled
				query := regexp.QuoteMeta(`INSERT INTO "users" ("username","email","phone") VALUES ($1,$2,$3) RETURNING "id"`)
				mock.ExpectBegin()
				mock.ExpectQuery(query).
					WithArgs(*user.Username, *user.Email, *user.Phone).
					WillReturnError(err)
				mock.ExpectRollback()
				return err
			},
			expectedError: sqlmock.ErrCancelled,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			mock, userRepository := createUserRepository()
			tt.mockFunction(mock, tt.user)
			err := userRepository.CreateUser(tt.user)
			if err != tt.expectedError {
				t.Errorf("Expected error: %v, got: %v", tt.expectedError, err)
			}
		})
	}
}
