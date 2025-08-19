package routers

import (
	"task4/controllers"
	"task4/middleware"

	"github.com/gin-gonic/gin"
)

func PostRoutersInit(r *gin.Engine) {
	postRouters := r.Group("/posts")
	postRouters.Use(middleware.AuthMiddleInit)
	{
		postRouters.POST("/create", controllers.PostController{}.CreatePost)
		postRouters.GET("/getPosts", controllers.PostController{}.GetPosts)
		postRouters.GET("/getpost", controllers.PostController{}.GetPost)
		postRouters.POST("/update", controllers.PostController{}.UpdatePost)
		postRouters.GET("/delete", controllers.PostController{}.DeletePost)
	}
}
