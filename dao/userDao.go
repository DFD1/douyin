package dao

import "log"

type User struct {
	Id       int64
	Name     string
	Password string
	Token    string
}

func (user User) TableName() string {
	return "users"
}

func QueryByToken(token string) (User, error) {
	user := User{}
	err := Db.Where("token = ?", token).First(&user).Error
	if err != nil {
		log.Println(err)
		return user, err
	}
	return user, nil
}

func InsertUser(user *User) bool {
	err := Db.Create(&user).Error
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
