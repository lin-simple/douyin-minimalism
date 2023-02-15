package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

type PublishListResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
}

func Publish(c *gin.Context) {
	// *注：title 未做请求处理
	token := c.PostForm("token")

	if _, existed := DemoUserLoginInfo[token]; !existed {
		c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "User doesn't exist."})
		return
	}

	file, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	fileName := filepath.Base(file.Filename)
	user := DemoUserLoginInfo[token]
	fileFinalName := fmt.Sprintf("%d_%s", user.Id, fileName) // 最后文件名格式为 user.id_filename
	dst := filepath.Join("./public/", fileFinalName)
	// fmt.Println("final file name: " + dst)

	if err := c.SaveUploadedFile(file, dst); err != nil { // 上传文件
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: fileFinalName + " upload success."})
}

func PublishList(c *gin.Context) {
	// token := c.Query("token")
	// user_id := c.Query("user_id")

	c.JSON(http.StatusOK, PublishListResponse{
		Response:  Response{StatusCode: 0},
		VideoList: DemoVideo,
	})
}
