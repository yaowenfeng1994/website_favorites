package handler

import (
	"fmt"
	"time"
	"db"
	"model"
)

type User struct {
	account  string
	nickname string
	mail     string
	time     int64
}


func Insert(){
	user := &User{
		account: "6107337181",
		nickname: "姚文锋",
		mail: "",
		time: time.Now().Unix(),
	}
	pool := db.InitMySQLPool("127.0.0.1", "website_favorites", "root", "123456", "utf8", 200, 100)
	fmt.Println(*user)
	//fmt.Println(pool)
	model.InsertUser(pool)
}