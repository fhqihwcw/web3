package routers

import (
	"task4/controllers"
	"task4/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutersInit(r *gin.Engine) {
	userRouters := r.Group("/users")
	userRouters.Use(middleware.AuthMiddleInit)
	{
		userRouters.POST("/register", controllers.UserController{}.Register)
		userRouters.POST("/login", controllers.UserController{}.Login)
	}
}
