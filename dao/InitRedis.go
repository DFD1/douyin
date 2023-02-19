package dao

import (
	"github.com/go-redis/redis"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
func InitRedisClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       1,  // use default DB
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	//err1 := rdb.Set("score", 100, time.Duration(100)*time.Second).Err()
	//if err != nil {
	//	fmt.Printf("set score failed,err:%v\n", err)
	//	return err1
	//}
	//val, err2 := rdb.Get("score").Result()
	//if err2 != nil {
	//	fmt.Printf("get score failed,err :%v\n", err)
	//	return err2
	//}
	//fmt.Println("score", val)
	return nil
}
