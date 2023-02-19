package dao

import (
	"time"
)

var VideoCount = 5 //每次获取视频流的数量

type Video struct {
	Id          int64  `json:"id"`
	PlayUrl     string `json:"play_url"`
	CoverUrl    string `json:"cover_url"`
	PublishTime time.Time
	Title       string `json:"title"`
}

func GetVideosByLastTime(lastTime time.Time) ([]Video, error) {
	videos := make([]Video, VideoCount)
	err := Db.Where("publish_time<?", lastTime).Order("publish_time desc").Limit(VideoCount).Find(&videos).Error
	return videos, err
}
