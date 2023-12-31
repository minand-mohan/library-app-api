package validator

import (
	"errors"
	"net/mail"

	"github.com/minand-mohan/library-app-api/api/users/dto"
	"github.com/minand-mohan/library-app-api/utils"
)

type UserValidator interface {
	ValidateUser(requestBody *dto.UserRequestBody) error
	ValidateUserQueryParams(queryParams *dto.UserQueryParams) error
	// ValidateUpdate(user *model.User) error
	// ValidateDelete(user *model.User) error
}

type UserValidatorImpl struct {
	logger *utils.AppLogger
}

func NewUserValidator(logger utils.AppLogger) UserValidator {
	return &UserValidatorImpl{
		logger: &logger,
	}
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (validator *UserValidatorImpl) ValidateUser(userReq *dto.UserRequestBody) error {
	validator.logger.Info("Validate user")
	if userReq.Username == "" {
		validator.logger.Error("Username is empty")
		return errors.New("Username is empty")
	}
	if userReq.Email == "" {
		validator.logger.Error("Email is empty")
		return errors.New("Email is empty")
	}
	if userReq.Phone == "" {
		validator.logger.Error("Phone is empty")
		return errors.New("Phone is empty")
	}
	if !isValidEmail(userReq.Email) {
		validator.logger.Error("Email is invalid")
		return errors.New("Email is invalid")
	}

	return nil
}

func (validator *UserValidatorImpl) ValidateUserQueryParams(queryParams *dto.UserQueryParams) error {
	if queryParams.Email != "" && !isValidEmail(queryParams.Email) {
		validator.logger.Error("Email is invalid")
		return errors.New("Email is invalid")
	}

	return nil
}
