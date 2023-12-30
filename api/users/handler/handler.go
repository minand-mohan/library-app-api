package handler

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/minand-mohan/library-app-api/api/response"
	"github.com/minand-mohan/library-app-api/api/users/dto"
	"github.com/minand-mohan/library-app-api/api/users/service"
	"github.com/minand-mohan/library-app-api/api/users/validator"
	"github.com/minand-mohan/library-app-api/utils"
)

type UserHandler struct {
	service   service.UserService
	validator validator.UserValidator
}

func NewUserHandler(service service.UserService, validator validator.UserValidator) *UserHandler {
	return &UserHandler{
		service:   service,
		validator: validator,
	}
}

func (handler *UserHandler) CreateUser(ctx *fiber.Ctx) error {
	log := utils.NewLogger()
	log.Info("Create user")
	// responseBody := response.HTTPResponse{
	// 	Code:    200,
	// 	Message: "Not implemented",
	// 	Content: map[string]interface{}{},
	// }
	// response.WriteHTTPResponse(c, 200, &responseBody)
	// return nil
	var userReq *dto.UserRequestBody
	err := json.Unmarshal(ctx.Request().Body(), &userReq)
	if err != nil {
		log.Error(fmt.Sprintf("Error while unmarshalling request body %v", err))
		responseBody := response.HTTPResponse{
			Code:    400,
			Message: "Bad request, invalid request body",
			Content: map[string]interface{}{},
		}
		err := response.WriteHTTPResponse(ctx, 400, &responseBody)
		if err != nil {
			log.Error(fmt.Sprintf("Error while writing response %v", err))
			return err
		}
		return nil
	}
	err = handler.validator.ValidateCreateUser(userReq)
	if err != nil {
		log.Error(fmt.Sprintf("Error while validating request body %v", err))
		responseBody := response.HTTPResponse{
			Code:    400,
			Message: "Bad request, invalid request body",
			Content: map[string]interface{}{},
		}
		err = response.WriteHTTPResponse(ctx, 400, &responseBody)
		if err != nil {
			log.Error(fmt.Sprintf("Error while writing response %v", err))
			return err
		}
		return nil
	}

	responseBody, err := handler.service.CreateUser(userReq)
	if err != nil {
		log.Error(fmt.Sprintf("UserHandler: Error while creating user %v", err))
		err = response.WriteHTTPResponse(ctx, responseBody.Code, responseBody)
		if err != nil {
			log.Error(fmt.Sprintf("Error while writing response %v", err))
			return err
		}
		return nil
	}

	response.WriteHTTPResponse(ctx, 200, responseBody)
	return nil

}
