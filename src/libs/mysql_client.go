package libs

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

func (p *SQLConnPool) Query(queryStr string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := p.SQLDB.Query(queryStr, args...)
	if err != nil {
		log.Println(err)
		return []map[string]interface{}{}, err
	}
	defer rows.Close()
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	rowsMap := make([]map[string]interface{}, 0, 10)
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		rowMap := make(map[string]interface{})
		for i, col := range values {
			if col != nil {
				rowMap[columns[i]] = string(col.([]byte))
			}
		}
		rowsMap = append(rowsMap, rowMap)
	}
	if err = rows.Err(); err != nil {
		return []map[string]interface{}{}, err
	}
	return rowsMap, nil
}

func (p *SQLConnPool) execute(sqlStr string, args ...interface{}) (sql.Result, error) {
	return p.SQLDB.Exec(sqlStr, args...)
}

func (p *SQLConnPool) Insert(insertStr string, args ...interface{}) (int64, error) {
	result, err := p.execute(insertStr, args...)
	if err != nil {
		return 0, err
	}
	LastId, err := result.LastInsertId()
	return LastId, err
}

// SQLConnTransaction is for transaction connection
type SQLConnTransaction struct {
	SQLTX *sql.Tx
}

// Begin transaction
func (p *SQLConnPool) Begin() (*SQLConnTransaction, error) {
	var oneSQLConnTransaction = &SQLConnTransaction{}
	var err error
	if pingErr := p.SQLDB.Ping(); pingErr == nil {
		oneSQLConnTransaction.SQLTX, err = p.SQLDB.Begin()
	}
	return oneSQLConnTransaction, err
}

// Rollback transaction
func (t *SQLConnTransaction) Rollback() error {
	return t.SQLTX.Rollback()
}

// Commit transaction
func (t *SQLConnTransaction) Commit() error {
	return t.SQLTX.Commit()
}

func (t *SQLConnTransaction) execute(sqlStr string, args ...interface{}) (sql.Result, error) {
	return t.SQLTX.Exec(sqlStr, args...)
}

func (t *SQLConnTransaction) Insert(insertStr string, args ...interface{}) (int64, error) {
	result, err := t.execute(insertStr, args...)
	if err != nil {
		return 0, err
	}
	LastId, err := result.LastInsertId()
	return LastId, err
}
