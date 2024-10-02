package repository

import (
	"regexp"
	"testing"

	"github.com/google/uuid"
	"github.com/minand-mohan/library-app-api/database/models"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestUpdateUser(t *testing.T) {
	//initialize variables
	test_email := "test@example.com"
	test_username := "test"
	test_phone := "1234567890"

	tc := []struct {
		name          string
		user          *models.User
		id            uuid.UUID
		mockFunction  func(mock sqlmock.Sqlmock, id uuid.UUID, user *models.User) error
		expectedError error
	}{
		{
			name: "User updated successfully",
			user: &models.User{
				Email:    &test_email,
				Phone:    &test_phone,
				Username: &test_username,
			},
			id: uuid.New(),
			mockFunction: func(mock sqlmock.Sqlmock, id uuid.UUID, user *models.User) error {
				query := regexp.QuoteMeta(`UPDATE "users" SET "username"=$1,"email"=$2,"phone"=$3 WHERE id = $4`)
				mock.ExpectBegin()
				mock.ExpectExec(query).
					WithArgs(*user.Username, *user.Email, *user.Phone, id).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
				return nil
			},
			expectedError: nil,
		},
		{
			name: "User update failed",
			user: &models.User{
				Email:    &test_email,
				Phone:    &test_phone,
				Username: &test_username,
			},
			id: uuid.New(),
			mockFunction: func(mock sqlmock.Sqlmock, id uuid.UUID, user *models.User) error {
				err := sqlmock.ErrCancelled
				query := regexp.QuoteMeta(`UPDATE "users" SET "username"=$1,"email"=$2,"phone"=$3 WHERE id = $4`)
				mock.ExpectBegin()
				mock.ExpectExec(query).
					WithArgs(*user.Username, *user.Email, *user.Phone, id).
					WillReturnError(err)
				mock.ExpectRollback()
				return err
			},
			expectedError: sqlmock.ErrCancelled,
		},
		{
			name: "User Partial Update",
			user: &models.User{
				Email: &test_email,
			},
			id: uuid.New(),
			mockFunction: func(mock sqlmock.Sqlmock, id uuid.UUID, user *models.User) error {
				query := regexp.QuoteMeta(`UPDATE "users" SET "email"=$1 WHERE id = $2`)
				mock.ExpectBegin()
				mock.ExpectExec(query).
					WithArgs(*user.Email, id).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
				return nil
			},
			expectedError: nil,
		},
	}
	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			mock, userRepository := createUserRepository()
			tt.mockFunction(mock, tt.id, tt.user)
			_, err := userRepository.UpdateByUserId(tt.id, tt.user)
			if err != tt.expectedError {
				t.Errorf("Expected error: %v, got: %v", tt.expectedError, err)
			}
		})
	}
}
