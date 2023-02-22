package controller

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	video_id := c.Query("video_id")
	action_type := c.Query("action_type")
	video_id_int64, _ := strconv.ParseInt(video_id, 10, 64)       //将string类型的video_id转换为int
	action_type_int64, _ := strconv.ParseInt(action_type, 10, 64) //将string类型的action_type转换为int
	user, err := dao.QueryByToken(token)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "用户未登录",
		})
	}
	if action_type_int64 == 1 {
		ok1 := dao.InsertLike(user.Id, video_id_int64, action_type_int64)
		ok2 := dao.LikeVideo(video_id_int64, user.Id, action_type_int64)
		if ok1 == true && ok2 == true {
			c.JSON(http.StatusOK, Response{
				StatusCode: 0,
			})
		} else {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "点赞失败"})
		}
	}
	if action_type_int64 == 2 {
		ok1 := dao.UpdateLike(user.Id, video_id_int64, action_type_int64)
		ok2 := dao.LikeVideo(video_id_int64, user.Id, action_type_int64)
		if ok1 == true && ok2 == true {
			c.JSON(http.StatusOK, Response{
				StatusCode: 0,
			})
		} else {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  "取消点赞失败",
			})
		}
	}

	//if _, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, Response{StatusCode: 0})
	//} else {
	//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	user_id := c.Query("user_id")
	token := c.Query("token")
	user_id_int64, _ := strconv.ParseInt(user_id, 10, 64)
	user, err := dao.QueryByToken(token)
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 0,
			},
			VideoList: DemoVideos,
		})
	}
	user_reponse := User{
		Id:            user.Id,
		Name:          user.Name,
		FollowCount:   3, //未开发
		FollowerCount: 4, //未开发
		IsFollow:      false,
	}
	video, _ := dao.GetAllVideoById(user_id_int64)
	var video_response []Video
	if len(video) == 0 {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 0,
			},
			VideoList: video_response,
		})
	} else {
		video_response = make([]Video, len(video))
		for i := 0; i < len(video); i++ {
			video_response[i].Id = video[i].Id
			video_response[i].Title = video[i].Title
			video_response[i].Author = user_reponse
			video_response[i].PlayUrl = video[i].PlayUrl
			video_response[i].CoverUrl = video[i].CoverUrl
			video_response[i].FavoriteCount = video[i].FavoriteCount
			video_response[i].IsFavorite = video[i].IsFavorite
			video_response[i].CommentCount = video[i].CommentCount
		}
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 0,
			},
			VideoList: video_response,
		})

	}
	//c.JSON(http.StatusOK, VideoListResponse{
	//	Response: Response{
	//		StatusCode: 0,
	//	},
	//	VideoList: DemoVideos,
	//})
}
