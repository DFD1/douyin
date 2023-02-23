package main

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/gin-gonic/gin"
)

func main() {
	//go service.RunMessageServer()
	dao.Init()
	dao.InitRedisClient()
	r := gin.Default()
	r.Static("/public", "./public")
	//r.StaticFS("/public", http.Dir("public"))

	//video_id := []int64{1, 2}
	//res := dao.QueryAllLikeVideoById(video_id)
	//fmt.Println(res)
	initRouter(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
