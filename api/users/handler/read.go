package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/minand-mohan/library-app-api/api/response"
	"github.com/minand-mohan/library-app-api/api/users/dto"
	"github.com/minand-mohan/library-app-api/utils"
)

func (handler *UserHandler) FindAllUsers(ctx *fiber.Ctx) error {
	log := utils.NewLogger()
	log.Info("Find all users")

	queryParams := new(dto.UserQueryParams)
	err := ctx.QueryParser(queryParams)
	if err != nil {
		log.Error(fmt.Sprintf("Error while parsing query params %v", err))
		responseBody := response.HTTPResponse{
			Code:    400,
			Message: "Bad request, invalid query params",
			Content: map[string]interface{}{},
		}
		err = response.WriteHTTPResponse(ctx, 400, &responseBody)
		if err != nil {
			log.Error(fmt.Sprintf("Error while writing response %v", err))
			return err
		}
		return nil
	}

	err = handler.validator.ValidateUserQueryParams(queryParams)
	if err != nil {
		log.Error(fmt.Sprintf("Error while validating query params %v", err))
		responseBody := response.HTTPResponse{
			Code:    400,
			Message: "Bad request, invalid query params",
			Content: map[string]interface{}{},
		}
		err = response.WriteHTTPResponse(ctx, 400, &responseBody)
		if err != nil {
			log.Error(fmt.Sprintf("Error while writing response %v", err))
			return err
		}
		return nil
	}

	responseBody, err := handler.service.FindAllUsers(queryParams)
	if err != nil {
		log.Error(fmt.Sprintf("UserHandler: Error while finding all users %v", err))
		err = response.WriteHTTPResponse(ctx, responseBody.Code, responseBody)
		if err != nil {
			log.Error(fmt.Sprintf("Error while writing response %v", err))
			return err
		}
		return nil
	}
	err = response.WriteHTTPResponse(ctx, 200, responseBody)
	if err != nil {
		log.Error(fmt.Sprintf("Error while writing response %v", err))
		return err
	}
	return nil
}

func (handler *UserHandler) FindByUserId(ctx *fiber.Ctx) error {
	log := utils.NewLogger()
	log.Info("Find user by id")
	id := ctx.Params("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		log.Error(fmt.Sprintf("Error while parsing uuid %v", err))
		responseBody := response.HTTPResponse{
			Code:    400,
			Message: "Bad request, invalid id",
			Content: map[string]interface{}{},
		}
		err = response.WriteHTTPResponse(ctx, 400, &responseBody)
		if err != nil {
			log.Error(fmt.Sprintf("Error while writing response %v", err))
			return err
		}
		return nil
	}
	responseBody, err := handler.service.FindByUserId(uuid)
	if err != nil {
		log.Error(fmt.Sprintf("UserHandler: Error while finding user by id %v", err))
		err = response.WriteHTTPResponse(ctx, responseBody.Code, responseBody)
		if err != nil {
			log.Error(fmt.Sprintf("Error while writing response %v", err))
			return err
		}
		return nil
	}
	err = response.WriteHTTPResponse(ctx, 200, responseBody)
	if err != nil {
		log.Error(fmt.Sprintf("Error while writing response %v", err))
		return err
	}
	return nil
}
