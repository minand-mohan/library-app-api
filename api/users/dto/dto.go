package dto

type UserRequestBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type UserQueryParams struct {
	Username string `query:"username"`
	Email    string `query:"email"`
}
