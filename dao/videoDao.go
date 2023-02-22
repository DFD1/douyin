package dao

import (
	"github.com/RaymondCode/simple-demo/config"
	"log"
	"time"
)

var VideoCount = config.VideoCount //每次最多获取视频流的数量

type Video struct {
	Id          int64     `json:"id"`
	AuthorId    int64     `json:"author_id"`
	PlayUrl     string    `json:"play_url"`
	CoverUrl    string    `json:"cover_url"`
	PublishTime time.Time `json:"publish_time"`
	Title       string    `json:"title"`
}

//获取离最近一次登录时间内投稿的视频
func GetVideosByLastTime(lastTime time.Time) ([]Video, error) {
	videos := make([]Video, VideoCount)
	err := Db.Where("publish_time < ?", lastTime).Order("publish_time desc").Limit(VideoCount).Find(&videos).Error
	return videos, err
}

//发布视频
func InsertVideo(authorid int64, playurl string, coverurl string, publishtime time.Time, title string) bool {
	video := Video{
		AuthorId:    authorid,
		PlayUrl:     playurl,
		CoverUrl:    coverurl,
		PublishTime: publishtime,
		Title:       title,
	}
	err := Db.Create(&video).Error
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

//根据发布视频用户id查询该用户发布的所有视频
func GetAllVideoById(author_id int64) ([]Video, error) {
	var videos []Video
	err := Db.Where("author_id", author_id).Find(&videos).Error
	return videos, err
}
