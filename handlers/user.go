package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync/atomic"
)

type UserRegisterResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserInfoResponse struct {
	Response
	User User `json:"user"`
}

var UserIdNum = int64(1)

// 用户注册
func UserRegister(c *gin.Context) { // 登录和注册共用该结构体
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if _, existed := DemoUserLoginInfo[token]; existed {
		c.JSON(http.StatusOK, UserRegisterResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User has already exist."},
		})
	} else {
		// 用户不存在，注册
		atomic.AddInt64(&UserIdNum, 1)
		newUser := User{
			Id:   UserIdNum,
			Name: username,
		}
		DemoUserLoginInfo[token] = newUser
		c.JSON(http.StatusOK, UserRegisterResponse{
			Response: Response{StatusCode: 0, StatusMsg: "User register success."},
			UserId:   UserIdNum,
			Token:    token,
		})
	}
	// fmt.Printf("user_id: %d, token: %s", UserIdNum, token)
}

// 用户登录
func UserLogin(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if user, existed := DemoUserLoginInfo[token]; existed {
		c.JSON(http.StatusOK, UserRegisterResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserRegisterResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist."},
		})
	}
}

// 用户信息
func UserInfo(c *gin.Context) {
	token := c.Query("token")

	if user, existed := DemoUserLoginInfo[token]; existed {
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: Response{StatusCode: 0},
			User: User{
				Id:            user.Id,
				Name:          user.Name,
				FollowerCount: user.FollowerCount,
				FollowCount:   user.FollowerCount,
				IsFollow:      user.IsFollow,
			},
		})
	} else {
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist."},
		})
	}
}
