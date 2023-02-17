package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"
)

var tmpChat = map[string][]Message{}

var msgIdSequence = int64(1)

type MessageListResponse struct {
	Response
	MessageList []Message `json:"message_list,omitempty"`
}

// 发送消息
func MessageAction(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	content := c.Query("content")

	if user, existed := DemoUserLoginInfo[token]; existed {
		userIdB, _ := strconv.Atoi(toUserId)
		tmpChatKey := getChatKey(user.Id, int64(userIdB))

		atomic.AddInt64(&msgIdSequence, 1)
		curMsg := Message{
			Id:         msgIdSequence,
			Content:    content,
			CreateTime: time.Now().Format(time.Kitchen),
		}
		// 处理消息情况
		if msg, existed := tmpChat[tmpChatKey]; existed {
			tmpChat[tmpChatKey] = append(msg, curMsg)
		} else {
			tmpChat[tmpChatKey] = []Message{curMsg}
		}

		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist."})
	}
}

func getChatKey(userIdA int64, userIdB int64) string {
	if userIdA > userIdB {
		return fmt.Sprintf("%d_%d", userIdB, userIdA)
	}
	return fmt.Sprintf("%d_%d", userIdA, userIdB)
}

// 聊天记录
func MessageChat(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")

	if user, existed := DemoUserLoginInfo[token]; existed {
		userIdB, _ := strconv.Atoi(toUserId)
		tmpChatKey := getChatKey(user.Id, int64(userIdB))

		c.JSON(http.StatusOK, MessageListResponse{
			Response:    Response{StatusCode: 0},
			MessageList: tmpChat[tmpChatKey],
		})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist."})
	}
}
