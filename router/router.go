package router

import (
	"douyin-minimalism/handlers"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.Static("/static", "./public")

	groupRouter := r.Group("/douyin")

	// basic apis
	groupRouter.GET("feed", handlers.Feed)
}
