package model

import (
	"fmt"
	"log"
	mysqlPool "db"
)


func InsertUser(mysqlPool *mysqlPool.SQLConnPool){
	//fmt.Println(user)
	lastuserid, err := mysqlPool.Insert("INSERT INTO user (`account`, `nickname`, `mail`, `create_time`)" +
		" VALUES( ?, ?, ?, ?)", "610234", "yaowenfeng", "610733719@qq.com", 1000000000)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(lastuserid)
	log.Fatal("create account success!")
}
