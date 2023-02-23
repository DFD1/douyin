package dao

import "log"

type Like struct {
	UserId  int64 `json:"user_id"`
	VideoId int64 `json:"video_id"`
	Cancel  int64 `json:"cancel"`
}

func (like Like) TableName() string {
	return "likes"
}

//当用户点赞时将点赞信息存入likes表中
func InsertLike(user_id int64, video_id int64, cancel int64) bool {
	favorite := Like{UserId: user_id, VideoId: video_id, Cancel: cancel}
	err := Db.Create(&favorite).Error
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

//当用户取消点赞时，将cancel的值更新为2
func UpdateLike(user_id int64, video_id int64, cancel int64) bool {
	err := Db.Table("likes").Where(&Like{UserId: user_id, VideoId: video_id, Cancel: 1}).Update("cancel", cancel).Error
	if err != nil {
		log.Panicln(err)
		return false
	}
	return true
}

//查询用户点赞的视频id的集合
func QueryLikeVideoId(user_id int64) []int64 {
	var res []int64
	Db.Table("likes").Where(&Like{UserId: user_id, Cancel: 1}).Select("video_id").Find(&res)
	return res
}
