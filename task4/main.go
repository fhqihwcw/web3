package main

import (
	"task4/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routers.UserRoutersInit(r)
	routers.PostRoutersInit(r)
	routers.CommentRoutersInit(r)
	r.Run()
}
