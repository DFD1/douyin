package controller

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")
	video_id := c.Query("video_id")
	comment_id := c.Query("comment_id")
	video_id_int64, _ := strconv.ParseInt(video_id, 10, 64)
	comment_id_int64, _ := strconv.ParseInt(comment_id, 10, 64)
	user, err := dao.QueryByToken(token)
	user_res := User{
		Id:            user.Id,
		Name:          user.Name,
		FollowerCount: 3,     //未开发
		FollowCount:   4,     //未开发
		IsFollow:      false, //未开发
	}
	if err == nil {
		if actionType == "1" {
			text := c.Query("comment_text")
			comment, ok := dao.InsertComment(user.Id, video_id_int64, text, time.Now(), 0)
			if ok == true {
				comment_res := Comment{
					Id:         comment.Id,
					User:       user_res,
					Content:    text,
					CreateDate: time.Now().Format("2006-01-02 15:04:05"),
				}
				c.JSON(http.StatusOK, CommentActionResponse{
					Response: Response{StatusCode: 0},
					Comment:  comment_res,
				})
			} else {
				c.JSON(http.StatusOK, CommentActionResponse{
					Response: Response{StatusCode: 1},
				})
			}
		}
		if actionType == "2" {
			ok := dao.DeleteComment(comment_id_int64)
			if ok == true {
				c.JSON(http.StatusOK, CommentActionResponse{
					Response: Response{
						StatusCode: 0,
					},
				})
			} else {
				c.JSON(http.StatusOK, CommentActionResponse{
					Response: Response{
						StatusCode: 1,
					},
				})
			}
		}

	}
	//if user, exist := usersLoginInfo[token]; exist {
	//	if actionType == "1" {
	//		text := c.Query("comment_text")
	//		c.JSON(http.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0},
	//			Comment: Comment{
	//				Id:         1,
	//				User:       user,
	//				Content:    text,
	//				CreateDate: "05-01",
	//			}})
	//		return
	//	}
	//	c.JSON(http.StatusOK, Response{StatusCode: 0})
	//} else {
	//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	token := c.Query("token")
	video_id := c.Query("video_id")
	video_id_int64, _ := strconv.ParseInt(video_id, 10, 64)
	user, err := dao.QueryByToken(token)
	if err != nil {
		c.JSON(http.StatusOK, CommentListResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
	user_res := User{
		Id:            user.Id,
		Name:          user.Name,
		FollowerCount: 3,     //未开发
		FollowCount:   4,     //未开发
		IsFollow:      false, //未开发
	}
	comments, ok := dao.GetCommentById(video_id_int64)
	if ok == true {
		comments_res := make([]Comment, len(comments))
		for i := 0; i < len(comments); i++ {
			comments_res[i].Id = comments[i].Id
			comments_res[i].User = user_res
			comments_res[i].Content = comments[i].CommentText
			comments_res[i].CreateDate = comments[i].CreateDate.Format("2006-01-02 15:04:05")
		}
		c.JSON(http.StatusOK, CommentListResponse{
			Response:    Response{StatusCode: 0},
			CommentList: comments_res,
		})
	} else {
		c.JSON(http.StatusOK, CommentListResponse{
			Response: Response{
				StatusCode: 1,
			},
		})
	}
}
