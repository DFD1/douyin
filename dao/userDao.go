package dao

import (
	"encoding/json"
	"log"
	"strconv"
)

type User struct {
	Id       int64
	Name     string
	Password string
	Token    string
}

func (user User) TableName() string {
	return "users"
}

//通过数据库Token查询用户
func QueryByToken(token string) (User, error) {
	user := User{}
	err := Db.Where("token = ?", token).First(&user).Error
	if err != nil {
		log.Println(err)
		return user, err
	}
	return user, nil
}

//将用户插入到mysql中
func InsertUser(user *User) bool {
	err := Db.Create(&user).Error
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func InserUser_redis(user *User) bool {
	key := "user:" + strconv.FormatInt(user.Id, 10)
	b, _ := json.Marshal(user)
	err := rdb.Set(key, b, 0).Err()
	if err != nil {
		log.Println(err)
		return false
	}
	//var res User
	//b1, _ := rdb.Get(key).Bytes()
	//json.Unmarshal(b1, &res)
	return true
}
