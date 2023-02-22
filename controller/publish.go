package controller

import (
	"bytes"
	"fmt"
	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	title := c.PostForm("title")
	user, error := dao.QueryByToken(token)
	if error != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "User doesn't exist",
		})
		return
	}
	//if _, exist := usersLoginInfo[token]; !exist {
	//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//	return
	//}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	//user := usersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := path.Join("./public", finalName)
	imageName, _ := GetImageFileName(saveFile)
	sqlPlayUrl := config.SqlUrl + finalName
	sqlCoverUrl := config.SqlUrl + imageName + ".png"
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	imageName_ := "./public/" + imageName
	GetSnapshot(saveFile, imageName_, 50)
	if err == nil {
		ok := dao.InsertVideo(user.Id, sqlPlayUrl, sqlCoverUrl, time.Now(), title)
		if ok != true {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  "uploaded failed",
			})
		} else {
			c.JSON(http.StatusOK, Response{
				StatusCode: 0,
				StatusMsg:  finalName + " uploaded successfully",
			})
		}

	}

}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token := c.Query("token")
	user_id := c.Query("user_id")
	user_id_int64, _ := strconv.ParseInt(user_id, 10, 64)
	video, _ := dao.GetAllVideoById(user_id_int64)
	user, _ := dao.QueryByToken(token)
	user_res := User{
		Id:            user.Id,
		Name:          user.Name,
		FollowerCount: 3,     //未开发
		FollowCount:   4,     //未开发
		IsFollow:      false, //未开发
	}
	l := len(video)
	video_res := make([]Video, l)
	for i := 0; i < l; i++ {
		video_res[i].Id = video[i].Id
		video_res[i].PlayUrl = video[i].PlayUrl
		video_res[i].Title = video[i].Title
		video_res[i].CoverUrl = video[i].CoverUrl
		video_res[i].FavoriteCount = video[i].FavoriteCount //未开发
		video_res[i].CommentCount = video[i].CommentCount   //未开发
		video_res[i].IsFavorite = video[i].IsFavorite       //未开发
		video_res[i].Author = user_res
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: video_res,
	})
}

func GetImageFileName(videoPath string) (string, error) {
	temp := strings.Split(videoPath, "/") // windows下使用\ ，linux下使用/
	videoName := temp[len(temp)-1]
	b := []byte(videoName)
	videoName = string(b[:len(b)-4])
	return videoName, nil
}

func GetSnapshot(videoPath, snapshotPath string, frameNum int) (snapshotName string, err error) {
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}

	err = imaging.Save(img, snapshotPath+".png")
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}

	names := strings.Split(snapshotPath, "\\")
	snapshotName = names[len(names)-1] + ".png"
	return
}
