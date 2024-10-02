package service

import (
	"github.com/google/uuid"
	"github.com/minand-mohan/library-app-api/database/models"
)

func generateRandomUser01() models.User {
	//initialize variables
	test_email := "test1@example.com"
	test_username := "test1"
	test_phone := "1234567890"
	test_id := uuid.New()
	return models.User{
		ID:       &test_id,
		Email:    &test_email,
		Username: &test_username,
		Phone:    &test_phone,
	}
}

func generateRandomUser02() models.User {
	//initialize variables
	test_email := "test2@example.com"
	test_username := "test2"
	test_phone := "1234567810"
	test_id := uuid.New()
	return models.User{
		ID:       &test_id,
		Email:    &test_email,
		Username: &test_username,
		Phone:    &test_phone,
	}
}
