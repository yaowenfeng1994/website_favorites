package lib

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

type SQLConnPool struct {
	DriverName     string
	DataSourceName string
	MaxOpenConns   int
	MaxIdleConns   int
	SQLDB          *sql.DB
}

func InitMySQLPool(host, database, user, password, charset string, maxOpenConns, maxIdleConns int) *SQLConnPool {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&autocommit=true", user, password, host, database,
		charset)
	db := &SQLConnPool{
		DriverName:     "mysql",
		DataSourceName: dataSourceName,
		MaxOpenConns:   maxOpenConns,
		MaxIdleConns:   maxIdleConns,
	}
	if err := db.open(); err != nil {
		log.Panicln("Init mysql pool failed.", err.Error())
	}
	return db
}

func (p *SQLConnPool) open() error {
	var err error
	p.SQLDB, err = sql.Open(p.DriverName, p.DataSourceName)
	if err != nil {
		return err
	}
	if err = p.SQLDB.Ping(); err != nil {
		return err
	}
	p.SQLDB.SetMaxOpenConns(p.MaxOpenConns)
	p.SQLDB.SetMaxIdleConns(p.MaxIdleConns)
	return nil
}

func (p *SQLConnPool) Close() error {
	return p.SQLDB.Close()
}

func (p *SQLConnPool) execute(sqlStr string, args ...interface{}) (sql.Result, error) {
	return p.SQLDB.Exec(sqlStr, args...)
}

func (p *SQLConnPool) Insert(insertStr string, args ...interface{}) (int64, error) {
	result, err := p.execute(insertStr, args...)
	if err != nil {
		return 0, err
	}
	lastid, err := result.LastInsertId()
	return lastid, err
}
