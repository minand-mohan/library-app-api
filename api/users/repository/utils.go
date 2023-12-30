package repository

import (
	"fmt"

	"github.com/minand-mohan/library-app-api/api/users/dto"
)

type UserDbQueries struct {
	Email    string
	Username string
}

func GenerateDbQueries(queryParams *dto.UserQueryParams) UserDbQueries {
	userDbQueries := UserDbQueries{}
	if queryParams.Email != "" {
		userDbQueries.Email = fmt.Sprintf("email = '%s'", queryParams.Email)
	}
	if queryParams.Username != "" {
		userDbQueries.Username = fmt.Sprintf("username ILIKE '%%%s%%'", queryParams.Username)
	}

	return userDbQueries
}
