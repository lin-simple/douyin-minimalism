package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

// 评论操作
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")

	if user, existed := DemoUserLoginInfo[token]; existed {
		if actionType == "1" {
			// 1 - 发布评论
			content := c.Query("comment_text")
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: Response{StatusCode: 0},
				Comment: Comment{
					Id:         1,
					User:       user,
					Content:    content,
					CreateDate: "2023-02-16",
				},
			})
			return
		}
		c.JSON(http.StatusOK, Response{StatusCode: 0})
		// delete comment.id??
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist."})
	}
}

// 评论列表
// All videos have same demo comment list.
func CommentList(c *gin.Context) {
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: DemoCommentList,
	})
}
