package model

import (
	"db"
)

var (
	Pool *db.SQLConnPool
)

func init() {
	Pool = db.InitMySQLPool("127.0.0.1", "website_favorites", "root", "123456", "utf8", 200, 100)
}

func InsertUser(account string, nickname string, mail string, t int64) (int64, error) {
	LastUserId, err := Pool.Insert("INSERT INTO user (`account`, `nickname`, `mail`, `create_time`)" +
		" VALUES( ?, ?, ?, ?)", account, nickname, mail, t)
	return LastUserId, err
}
