package routers

import (
	"task4/controllers"
	"task4/middleware"

	"github.com/gin-gonic/gin"
)

func CommentRoutersInit(r *gin.Engine) {
	commentRouters := r.Group("/comments")
	commentRouters.Use(middleware.AuthMiddleInit)
	{
		commentRouters.POST("/add", controllers.CommentController{}.CreateComment)
		commentRouters.GET("/get", controllers.CommentController{}.GetComments)
	}
}
