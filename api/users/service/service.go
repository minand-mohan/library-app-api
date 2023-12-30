package service

import (
	"github.com/google/uuid"
	"github.com/minand-mohan/library-app-api/api/response"
	"github.com/minand-mohan/library-app-api/api/users/dto"
	"github.com/minand-mohan/library-app-api/api/users/repository"
	"github.com/minand-mohan/library-app-api/utils"
)

type UserService interface {
	CreateUser(userReqBody *dto.UserRequestBody) (*response.HTTPResponse, error)
	FindAllUsers(queryParams *dto.UserQueryParams) (*response.HTTPResponse, error)
	FindByUserId(id uuid.UUID) (*response.HTTPResponse, error)
	UpdateByUserId(id uuid.UUID, userReqBody *dto.UserRequestBody) (*response.HTTPResponse, error)
	DeleteByUserId(id uuid.UUID) (*response.HTTPResponse, error)
}

type UserServiceImpl struct {
	repo   repository.UserRepository
	logger *utils.AppLogger
}

func NewUserService(repo repository.UserRepository, logger utils.AppLogger) UserService {
	return &UserServiceImpl{
		repo:   repo,
		logger: &logger,
	}
}
