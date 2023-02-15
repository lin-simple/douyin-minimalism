package router

import (
	"douyin-minimalism/handlers"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.Static("/static", "./public")

	groupRouter := r.Group("/douyin")

	// basic apis
	groupRouter.GET("/feed", handlers.Feed)
	groupRouter.POST("/user/register/", handlers.UserRegister)
	groupRouter.POST("/user/login/", handlers.UserLogin)
	groupRouter.GET("/user/", handlers.UserInfo)
}
