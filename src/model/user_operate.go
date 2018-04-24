package model

import (
	"fmt"
	"log"
	mysqlPool "db"
)


func InsertUser(mysqlPool *mysqlPool.SQLConnPool, user map[string]string){
	fmt.Println(user)
	lastuserid, err := mysqlPool.Insert("INSERT INTO user (`account`, `nickname`, `mail`, `create_time`)" +
		" VALUES( ?, ?, ?, ?)", user["account"], user["nickname"], user["mail"], user["create_time"])
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(lastuserid)
	log.Fatal("create account success!")
}
