package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lovelyrrg51/go_backend/app/controllers"
	"github.com/lovelyrrg51/go_backend/app/database"
	"github.com/lovelyrrg51/go_backend/app/middlewares"
	"github.com/lovelyrrg51/go_backend/app/repositories"
	"github.com/lovelyrrg51/go_backend/app/services"
)

func UserRouter(userRouter *gin.RouterGroup) {
	// Initialize
	newUserRepo := repositories.NewUserRepository(database.DB)
	newUserService := services.NewUserService(newUserRepo)
	newUserCtrl := controllers.NewUserController(newUserService)
	jwtMiddleware := middlewares.NewJWTMiddleware(newUserRepo)

	/* Auth API Routes */
	userRouter.Use(jwtMiddleware.Verify())
	{
		userRouter.GET("/profile", newUserCtrl.GetProfile())
	}
}
