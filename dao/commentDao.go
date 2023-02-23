package dao

import (
	"log"
	"time"
)

type Comment struct {
	Id          int64     `json:"id"`
	UserId      int64     `json:"user_id"`
	VideoId     int64     `json:"video_id"`
	CommentText string    `json:"comment_text"`
	CreateDate  time.Time `json:"create_date"`
	Cancel      int64     `json:"cancel"`
}

//插入一条评论
func InsertComment(user_id int64, video_id int64, comment_text string, create_date time.Time, cancel int64) (Comment, bool) {
	comment := Comment{
		UserId:      user_id,
		VideoId:     video_id,
		CommentText: comment_text,
		CreateDate:  create_date,
		Cancel:      cancel,
	}
	err := Db.Table("comments").Create(&comment).Error
	if err != nil {
		log.Println(err)
		return Comment{}, false
	}
	return comment, true
}

//删除一条评论
func DeleteComment(comment_id int64) bool {
	err := Db.Delete(&Comment{}, comment_id).Error
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

//查询视频的所有评论，按发布时间倒序
func GetCommentById(video_id int64) ([]Comment, bool) {
	var comments []Comment
	err := Db.Where("video_id", video_id).Order("create_date desc").Find(&comments).Error
	if err != nil {
		log.Println(err)
		return comments, false
	}
	return comments, true
}
