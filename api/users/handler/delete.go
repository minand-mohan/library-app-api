package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/minand-mohan/library-app-api/api/response"
	"github.com/minand-mohan/library-app-api/utils"
)

func (handler *UserHandler) DeleteByUserId(ctx *fiber.Ctx) error {
	log := utils.NewLogger()
	log.Info("Delete user by id")
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
	responseBody, err := handler.service.DeleteByUserId(uuid)
	if err != nil {
		log.Error(fmt.Sprintf("UserHandler: Error while deleting user by id %v", err))
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
