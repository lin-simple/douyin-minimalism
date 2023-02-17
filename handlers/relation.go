package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list,omitempty"`
}

// 关注操作
func Follow(c *gin.Context) {
	token := c.Query("token")

	if _, existed := DemoUserLoginInfo[token]; existed {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist."})
	}
}

// 关注列表
func FollowList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{StatusCode: 0},
		UserList: DemoUserList,
	})
}

// 粉丝列表
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{StatusCode: 0},
		UserList: DemoUserList,
	})
}

// 好友列表
func FriendList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{StatusCode: 0},
		UserList: DemoUserList,
	})
}
