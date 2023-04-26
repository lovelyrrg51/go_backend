package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lovelyrrg51/go_backend/app/common"
	"github.com/lovelyrrg51/go_backend/app/dtos"
	"github.com/lovelyrrg51/go_backend/app/services"
	"github.com/lovelyrrg51/go_backend/app/utils"
	"gopkg.in/validator.v2"
)

type AuthController struct {
	userService services.UserService
}

func NewAuthController(userService services.UserService) AuthController {
	return AuthController{
		userService: userService,
	}
}

func (aCtrl AuthController) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check Request
		var req dtos.AuthRegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.WriteResponseError(c, common.NewBadRequestError(err.Error()))
			return
		}

		// Check Validate
		if err := validator.Validate(req); err != nil {
			utils.WriteResponseError(c, common.NewBadRequestError(err.Error()))
			return
		}

		// Register User
		res, err := aCtrl.userService.Register(req)
		if err != nil {
			utils.WriteResponseError(c, err)
		}

		utils.WriteResponse(c, http.StatusOK, res)
	}
}

func (aCtrl AuthController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check Request
		var req dtos.AuthLoginRequest
		if err := c.BindJSON(&req); err != nil {
			utils.WriteResponseError(c, common.NewBadRequestError(err.Error()))
			return
		}

		// Check Validate
		if err := validator.Validate(req); err != nil {
			utils.WriteResponseError(c, common.NewBadRequestError(err.Error()))
			return
		}

		// Login User
		res, err := aCtrl.userService.Login(req)
		if err != nil {
			utils.WriteResponseError(c, err)
		}

		utils.WriteResponse(c, http.StatusOK, res)
	}
}
