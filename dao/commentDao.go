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
