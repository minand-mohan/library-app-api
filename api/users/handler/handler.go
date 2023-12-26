package handler

import (
	"github.com/minand-mohan/library-app-api/api/users/service"
	"github.com/minand-mohan/library-app-api/api/users/validator"
)

type UserHandler struct {
	service   service.UserService
	validator validator.UserValidator
}
