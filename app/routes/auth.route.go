package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lovelyrrg51/go_backend/app/controllers"
	"github.com/lovelyrrg51/go_backend/app/database"
	"github.com/lovelyrrg51/go_backend/app/repositories"
	"github.com/lovelyrrg51/go_backend/app/services"
)

func AuthRouter(authRouter *gin.RouterGroup) {
	// Initialize
	newUserRepo := repositories.NewUserRepository(database.DB)
	newUserService := services.NewUserService(newUserRepo)
	newAuthCtrl := controllers.NewAuthController(newUserService)

	/* Auth API Routes */
	// Register API
	authRouter.POST("/register", newAuthCtrl.Register())
	// Login API
	authRouter.POST("/login", newAuthCtrl.Login())
}
