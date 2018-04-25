package handler

import (
	"db"
	"model"
	"time"
	"strconv"
)

//type User struct {
//	account  string
//	nickname string
//	mail     string
//	time     int64
//}


func Insert(){
	//user := &User{
	//	account: "6107337181",
	//	nickname: "姚文锋",
	//	mail: "",
	//	time: time.Now().Unix(),
	//}
	var user map[string]string
	user = make(map[string]string)
	user["account"] = "143215"
	user["nickname"] = "14325"
	user["mail"] = ""
	user["create_time"] = strconv.FormatInt(time.Now().Unix(), 10)
	pool := db.InitMySQLPool("127.0.0.1", "website_favorites", "root", "1q2w3e", "utf8", 200, 100)
	model.InsertUser(pool, user)
}