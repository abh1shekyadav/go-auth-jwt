package routes

import (
	controller "github.com/abh1shekyadav/go-auth-jwt/controllers"
	"github.com/abh1shekyadav/go-auth-jwt/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
