package repository

import (
	"regexp"
	"testing"

	"github.com/google/uuid"
	"github.com/minand-mohan/library-app-api/api/users/dto"
	"github.com/minand-mohan/library-app-api/database/models"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestFindAllUsers(t *testing.T) {

	user1 := generateRandomUser01()
	user2 := generateRandomUser02()

	tc := []struct {
		name          string
		params        *dto.UserQueryParams
		mockFunction  func(mock sqlmock.Sqlmock, params dto.UserQueryParams) error
		expectedError error
		expectedList  []models.User
	}{
		{
			name:   "Find all users successfully",
			params: &dto.UserQueryParams{},
			mockFunction: func(mock sqlmock.Sqlmock, params dto.UserQueryParams) error {
				query := regexp.QuoteMeta(`SELECT * FROM "users"`)
				mock.ExpectQuery(query).
					WillReturnRows(sqlmock.NewRows([]string{"id", "username", "email", "phone"}).
						AddRow(user1.ID, user1.Username, user1.Email, user1.Phone).
						AddRow(user2.ID, user2.Username, user2.Email, user2.Phone))
				return nil
			},
			expectedError: nil,
			expectedList:  []models.User{user1, user2},
		},
		{
			name:   "Find no users",
			params: &dto.UserQueryParams{Username: "test 1234"},
			mockFunction: func(mock sqlmock.Sqlmock, params dto.UserQueryParams) error {
				query := regexp.QuoteMeta(`SELECT * FROM "users" WHERE username ILIKE '%test 1234%`)
				mock.ExpectQuery(query).
					WillReturnRows(sqlmock.NewRows([]string{"id", "username", "email", "phone"}))
				return nil
			},
			expectedError: nil,
			expectedList:  []models.User{},
		},
		{
			name:   "Find all users with error",
			params: &dto.UserQueryParams{Email: "test"},
			mockFunction: func(mock sqlmock.Sqlmock, params dto.UserQueryParams) error {
				err := sqlmock.ErrCancelled
				query := regexp.QuoteMeta(`SELECT * FROM "users" WHERE email = 'test'`)
				mock.ExpectQuery(query).
					WillReturnError(err)
				return err
			},
			expectedError: sqlmock.ErrCancelled,
			expectedList:  []models.User{},
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			mock, userRepository := createUserRepository()
			tt.mockFunction(mock, *tt.params)
			users, err := userRepository.FindAllUsers(tt.params)
			if err != tt.expectedError {
				t.Errorf("Expected error: %v, got: %v", tt.expectedError, err)
			}
			if len(users) != len(tt.expectedList) {
				t.Errorf("Expected list length: %v, got: %v", len(tt.expectedList), len(users))
			}
		})
	}
}

func TestFindUserByID(t *testing.T) {

	user := generateRandomUser01()

	tc := []struct {
		name          string
		id            uuid.UUID
		mockFunction  func(mock sqlmock.Sqlmock, id string) error
		expectedError error
		expectedUser  *models.User
	}{
		{
			name: "Find user by id successfully",
			id:   *user.ID,
			mockFunction: func(mock sqlmock.Sqlmock, id string) error {
				query := regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)
				mock.ExpectQuery(query).
					WithArgs(id).
					WillReturnRows(sqlmock.NewRows([]string{"id", "username", "email", "phone"}).
						AddRow(user.ID, user.Username, user.Email, user.Phone))
				return nil
			},
			expectedError: nil,
			expectedUser:  &user,
		},
		{
			name: "Find user by id with error",
			id:   *user.ID,
			mockFunction: func(mock sqlmock.Sqlmock, id string) error {
				err := sqlmock.ErrCancelled
				query := regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)
				mock.ExpectQuery(query).
					WithArgs(id).
					WillReturnError(err)
				return err
			},
			expectedError: sqlmock.ErrCancelled,
			expectedUser:  nil,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			mock, userRepository := createUserRepository()
			tt.mockFunction(mock, tt.id.String())
			user, err := userRepository.FindByUserId(tt.id)
			if err != tt.expectedError {
				t.Errorf("Expected error: %v, got: %v", tt.expectedError, err)
			}
			if user != nil && *user != *tt.expectedUser {
				t.Errorf("Expected user: %v, got: %v", tt.expectedUser, user)
			}
		})
	}
}

func TestFindByEmailOrUsernameOrPhone(t *testing.T) {

	outputUser := generateRandomUser01()
	test_username := "user"
	test_phone := "234567890"
	test_email := "user@example.com"

	tc := []struct {
		name          string
		inputUser     *models.User
		mockFunction  func(mock sqlmock.Sqlmock, user *models.User) error
		expectedError error
		expectedUser  *models.User
	}{
		{
			name: "Find user by email successfully",
			inputUser: &models.User{
				Email:    outputUser.Email,
				Username: &test_username,
				Phone:    &test_phone,
			},
			mockFunction: func(mock sqlmock.Sqlmock, user *models.User) error {
				query := regexp.QuoteMeta(`SELECT * FROM "users" WHERE email = $1 OR username = $2 OR phone = $3 ORDER BY "users"."id" LIMIT 1`)
				mock.ExpectQuery(query).
					WithArgs(user.Email, user.Username, user.Phone).
					WillReturnRows(sqlmock.NewRows([]string{"id", "username", "email", "phone"}).
						AddRow(outputUser.ID, outputUser.Username, outputUser.Email, outputUser.Phone))
				return nil
			},
			expectedError: nil,
			expectedUser:  &outputUser,
		},
		{
			name: "Find user by username successfully",
			inputUser: &models.User{
				Email:    &test_email,
				Username: outputUser.Username,
				Phone:    &test_phone,
			},
			mockFunction: func(mock sqlmock.Sqlmock, user *models.User) error {
				query := regexp.QuoteMeta(`SELECT * FROM "users" WHERE email = $1 OR username = $2 OR phone = $3 ORDER BY "users"."id" LIMIT 1`)
				mock.ExpectQuery(query).
					WithArgs(user.Email, user.Username, user.Phone).
					WillReturnRows(sqlmock.NewRows([]string{"id", "username", "email", "phone"}).
						AddRow(outputUser.ID, outputUser.Username, outputUser.Email, outputUser.Phone))
				return nil
			},
			expectedError: nil,
			expectedUser:  &outputUser,
		},
		{
			name: "Find user by phone successfully",
			inputUser: &models.User{
				Email:    &test_email,
				Username: &test_username,
				Phone:    outputUser.Phone,
			},
			mockFunction: func(mock sqlmock.Sqlmock, user *models.User) error {
				query := regexp.QuoteMeta(`SELECT * FROM "users" WHERE email = $1 OR username = $2 OR phone = $3 ORDER BY "users"."id" LIMIT 1`)
				mock.ExpectQuery(query).
					WithArgs(user.Email, user.Username, user.Phone).
					WillReturnRows(sqlmock.NewRows([]string{"id", "username", "email", "phone"}).
						AddRow(outputUser.ID, outputUser.Username, outputUser.Email, outputUser.Phone))
				return nil
			},
			expectedError: nil,
			expectedUser:  &outputUser,
		},
		{
			name: "Find user by email with error",
			inputUser: &models.User{
				Email:    outputUser.Email,
				Username: &test_username,
				Phone:    &test_phone,
			},
			mockFunction: func(mock sqlmock.Sqlmock, user *models.User) error {
				err := sqlmock.ErrCancelled
				query := regexp.QuoteMeta(`SELECT * FROM "users" WHERE email = $1 OR username = $2 OR phone = $3 ORDER BY "users"."id" LIMIT 1`)
				mock.ExpectQuery(query).
					WithArgs(user.Email, user.Username, user.Phone).
					WillReturnError(err)
				return err
			},
			expectedError: sqlmock.ErrCancelled,
			expectedUser:  nil,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			mock, userRepository := createUserRepository()
			tt.mockFunction(mock, tt.inputUser)
			user, err := userRepository.FindByEmailOrUsernameOrPhone(*tt.inputUser.Email, *tt.inputUser.Username, *tt.inputUser.Phone)
			if err != tt.expectedError {
				t.Errorf("Expected error: %v, got: %v", tt.expectedError, err)
			}
			if user != nil && *user != *tt.expectedUser {
				t.Errorf("Expected user: %v, got: %v", tt.expectedUser, user)
			}
		})
	}
}
