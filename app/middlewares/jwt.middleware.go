package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lovelyrrg51/go_backend/app/common"
	"github.com/lovelyrrg51/go_backend/app/repositories"
	"github.com/lovelyrrg51/go_backend/app/utils"
)

type JWTMiddleware struct {
	repo repositories.UserRepository
}

func NewJWTMiddleware(repo repositories.UserRepository) JWTMiddleware {
	return JWTMiddleware{
		repo: repo,
	}
}

func (m JWTMiddleware) Verify() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get auth token from Header
		authenticationHeader := c.GetHeader("Authorization")
		if authenticationHeader == "" {
			utils.WriteResponseError(c, common.NewUnauthenticatedError("Unauthorized"))
			return
		}

		splitAuthHeader := strings.Split(authenticationHeader, " ")
		if len(splitAuthHeader) != 2 {
			utils.WriteResponseError(c, common.NewUnauthenticatedError("Missing or invalid authorization header"))
			return
		}
		authToken := splitAuthHeader[1]

		// Verify Auth Token
		userId, err := utils.VerifyStandardJWTToken(authToken)
		if err != nil {
			utils.WriteResponseError(c, common.NewUnauthenticatedError("Invalid or expired authorization token"))
			return
		}

		// Get User from ID
		// existedUser, userErr := m.repo.FindById(*userId)
		// if userErr != nil {
		// 	utils.WriteResponseError(c, userErr)
		// 	return
		// }

		// Set User && ID
		// c.Set("user", existedUser)
		c.Set("userId", userId.String())

		// Next
		c.Next()
	}
}
