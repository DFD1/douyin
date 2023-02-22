package dao

import (
	"github.com/RaymondCode/simple-demo/config"
	"gorm.io/gorm"
	"log"
	"time"
)

var VideoCount = config.VideoCount //每次最多获取视频流的数量

type Video struct {
	Id            int64     `json:"id"`
	AuthorId      int64     `json:"author_id"`
	PlayUrl       string    `json:"play_url"`
	CoverUrl      string    `json:"cover_url"`
	PublishTime   time.Time `json:"publish_time"`
	Title         string    `json:"title"`
	FavoriteCount int64     `json:"favorite_count"`
	CommentCount  int64     `json:"comment_count"`
	IsFavorite    bool      `json:"is_favorite"`
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
		AuthorId:      authorid,
		PlayUrl:       playurl,
		CoverUrl:      coverurl,
		PublishTime:   publishtime,
		Title:         title,
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
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

//用户点赞时，将favorite_count加1,并将is_favorite改为true
//用户取消赞时，将favorite_count减1，并将is_favorite改为false
func LikeVideo(video_id int64, author_id int64, action_type int64) bool {
	if action_type == 1 {
		err1 := Db.Table("videos").Where(&Video{Id: video_id, AuthorId: author_id}).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
		err2 := Db.Table("videos").Where(&Video{Id: video_id, AuthorId: author_id}).UpdateColumn("is_favorite", true).Error
		if err1 != nil || err2 != nil {
			return false
		} else {
			return true
		}
	}
	if action_type == 2 {
		err1 := Db.Table("videos").Where(&Video{Id: video_id, AuthorId: author_id}).UpdateColumn("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
		err2 := Db.Table("videos").Where(&Video{Id: video_id, AuthorId: author_id}).UpdateColumn("is_favorite", false).Error
		if err1 != nil || err2 != nil {
			return false
		} else {
			return true
		}
	}
	return false
}

//查询视频的点赞数量
func QueryVideoFavoriteCount(video_id int64) int64 {
	var res int64
	Db.Table("videos").Where("id", video_id).Select("favorite_count").Scan(&res)
	return res
}
