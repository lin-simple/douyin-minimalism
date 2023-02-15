package main

import (
	"douyin-minimalism/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.InitRouter(r)

	err := r.Run()
	if err != nil {
		return
	}
}
