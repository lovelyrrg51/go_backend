package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lovelyrrg51/go_backend/app/services"
	"github.com/lovelyrrg51/go_backend/app/utils"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return UserController{
		userService: userService,
	}
}

func (uCtrl UserController) GetProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get User from Request
		userId, existed := c.Get("userId")
		if !existed {
			c.AbortWithStatusJSON(401, gin.H{"message": "User hasn't logged in yet"})
			return
		}

		// Get Profile
		res, err := uCtrl.userService.GetProfile(uuid.MustParse(userId.(string)))
		if err != nil {
			utils.WriteResponseError(c, err)
		}

		utils.WriteResponse(c, http.StatusOK, res)
	}
}
