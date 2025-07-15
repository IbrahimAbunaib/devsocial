package routes

import (
	"github.com/gin-gonic/gin"
	"devsocial-backend/controllers"
	"github.com/net/http"
)

auth := router.Group("/api/auth")
{
	auth.POST("/signup", controllers.Signup)
	auth.POST("/login", controllers.Login)
}
