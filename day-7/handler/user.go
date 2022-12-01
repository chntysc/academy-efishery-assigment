package handler

import (
	"mymodule/usecase"

	"mymodule/entity/response"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

func NewUserHandler(usercase *usecase.UserUseCase) *UserHandler{
	return &UserHandler{userUseCase: usercase}
}

func (h UserHandler) CreateUser(ctx *fiber.Ctx) error {
	userRequest := response.CreateUserRequest{}
	if err := ctx.BodyParser(&userRequest);err!=nil{
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:		fiber.StatusBadRequest,
			Message:	"invalid req body",
			Data:		nil,
		})
	}
	if err :=h.userUseCase.CreateUser(userRequest);err != nil{
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:		fiber.StatusOK,
			Message:	"failed to create data",
			Data:		nil,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Code:		fiber.StatusOK,
		Message:	"successfully create data",
		Data:		nil,
	})
}

func (h UserHandler) GetList(ctx *fiber.Ctx) error{
	users, err := h.userUseCase.GetList();

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:		fiber.StatusBadRequest,
			Message:	"get list users failed",
			Data:		nil,
		})
	}

	if len(users)<=0 {
		return ctx.Status(fiber.StatusNoContent).JSON(response.BaseResponse{
			Code:		fiber.StatusNoContent,
			Message:	"no user found",
			Data:		nil,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Code:		fiber.StatusOK,
		Message:	"successfully get all users",
		Data:		users,
	})


}



