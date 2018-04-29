package model

import (
	"libs"
)

var (
	Pool *libs.SQLConnPool
)

func init() {
	Pool = libs.InitMySQLPool("127.0.0.1", "website_favorites", "root", "1q2w3e", "utf8", 200, 100)
}

//func InsertUser(account string, nickname string, mail string, password string, t int64) (int64, error) {
//	LastUserId, err := Pool.Insert("INSERT INTO user (`account`, `nickname`, `mail`, `create_time`)" +
//		" VALUES( ?, ?, ?, ?)", account, nickname, mail, t)
//	if LastUserId != 0 {
//		fmt.Println(password, LastUserId)
//		_, err1 := Pool.Insert("INSERT INTO password (`password`, `user_id`, `create_time`)" +
//			" VALUES( ?, ?, ?)", password, LastUserId, t)
//		return LastUserId, err1
//	} else {
//		return LastUserId, err
//	}
//}

func InsertUser(account string, nickname string, mail string, password string, t int64) (int64, error) {
	Tx, err := Pool.Begin()
	if err != nil {
		return 0, err
	} else {
		LastUserId, err := Tx.Insert("INSERT INTO user (`account`, `nickname`, `mail`, `create_time`)" +
			" VALUES( ?, ?, ?, ?)", account, nickname, mail, t)
		if err != nil {
			Tx.Rollback()
			return 0, err
		} else {
			_, err1 := Tx.Insert("INSERT INTO password (`password`, `user_id`, `create_time`)" +
				" VALUES( ?, ?, ?)", password, LastUserId, t)
			if err1 != nil {
				Tx.Rollback()
				return 0, err1
			} else {
				err2 := Tx.Commit()
				return LastUserId, err2
			}
		}
	}
}

func QueryUserPassword(account string) (string, error) {

}
