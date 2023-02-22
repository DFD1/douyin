package controller

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

var VideoCount int //每次获取视频流的数量

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	token := c.Query("token")
	log.Println(token)

	usr, _ := dao.QueryByToken(token)
	user := User{
		Id:            usr.Id,
		Name:          usr.Name,
		FollowerCount: 3,     //社交功能未开发
		FollowCount:   4,     //社交功能未开发
		IsFollow:      false, //社交功能未开发
	}
	var latest_time time.Time
	latest_time = time.Now()
	video, err := dao.GetVideosByLastTime(latest_time)
	if err != nil {
		log.Println("没有视频投稿")
	}
	if len(video) >= 30 {
		VideoCount = 30
	} else {
		VideoCount = len(video)
	}

	video_res := make([]Video, VideoCount)
	for i := 0; i < VideoCount; i++ {
		video_res[i].Id = video[i].Id
		video_res[i].PlayUrl = video[i].PlayUrl
		video_res[i].Author = user
		video_res[i].CoverUrl = video[i].CoverUrl
		video_res[i].IsFavorite = false //未开发
		video_res[i].CommentCount = 0   //未开发
		video_res[i].FavoriteCount = 0  //未开发
		video_res[i].Title = video[i].Title
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: video_res,
		NextTime:  time.Now().Unix(),
	})
}
