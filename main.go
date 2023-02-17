package main

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/gin-gonic/gin"
)

func main() {
	//go service.RunMessageServer()
	dao.Init()
	r := gin.Default()
	//user, err := dao.QueryByToken("32423")
	//fmt.Println("%v\n", user)
	//if err != nil {
	//	fmt.Println(err)
	//}
	initRouter(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
