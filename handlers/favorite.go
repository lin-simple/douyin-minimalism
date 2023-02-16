package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Just check token if vaild.
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")

	if _, existed := DemoUserLoginInfo[token]; existed {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist."})
	}
}

// All users have same favorite video list
func FavoriteList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response:  Response{StatusCode: 0},
		VideoList: DemoVideo,
	})
}
