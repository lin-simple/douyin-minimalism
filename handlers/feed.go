package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	Response
	NextTime  int64   `json:"next_time"`
	VideoList []Video `json:"video_list,omitempty"`
}

func Feed(c *gin.Context) {
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		NextTime:  time.Now().Unix(),
		VideoList: DemoVideo,
	})
}
