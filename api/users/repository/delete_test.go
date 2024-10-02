package repository

import (
	"regexp"
	"testing"

	"github.com/google/uuid"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestDeleteUser(t *testing.T) {

	tc := []struct {
		name          string
		id            uuid.UUID
		mockFunction  func(mock sqlmock.Sqlmock, id uuid.UUID) error
		expectedError error
	}{
		{
			name: "Delete User by id successfully",
			id:   uuid.New(),
			mockFunction: func(mock sqlmock.Sqlmock, id uuid.UUID) error {
				query := regexp.QuoteMeta(`DELETE FROM "users" WHERE "users"."id" = $1`)
				mock.ExpectBegin()
				mock.ExpectExec(query).
					WithArgs(id.String()).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
				return nil
			},
			expectedError: nil,
		},
		{
			name: "Delete User by id with error",
			id:   uuid.New(),
			mockFunction: func(mock sqlmock.Sqlmock, id uuid.UUID) error {
				err := sqlmock.ErrCancelled
				query := regexp.QuoteMeta(`DELETE FROM "users" WHERE "users"."id" = $1`)
				mock.ExpectBegin()
				mock.ExpectExec(query).
					WithArgs(id.String()).
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
			tt.mockFunction(mock, tt.id)
			err := userRepository.DeleteByUserId(tt.id)
			if err != tt.expectedError {
				t.Errorf("Expected error: %v, got: %v", tt.expectedError, err)
			}
		})
	}
}
