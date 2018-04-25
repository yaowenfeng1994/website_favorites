package model

import (
	mysqlPool "lib"
)

func InsertUser(mysqlPool *mysqlPool.SQLConnPool, account string, nickname string, mail string, t int64) (int64, error) {
	LastUserId, err := mysqlPool.Insert("INSERT INTO user (`account`, `nickname`, `mail`, `create_time`)" +
		" VALUES( ?, ?, ?, ?)", account, nickname, mail, t)
	return LastUserId, err
}
