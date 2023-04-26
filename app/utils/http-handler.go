package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/lovelyrrg51/go_backend/app/common"
)

func WriteResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(code, data)
}

func WriteResponseError(c *gin.Context, err *common.AppError) {
	c.Error(errors.New(err.Message))
	c.AbortWithStatusJSON(err.Code, err.Error())
}
