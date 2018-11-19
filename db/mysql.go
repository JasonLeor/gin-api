package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

/**
数据库连接
*/
var SqlDB *sql.DB

func init() {
	SqlDB, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3307)/gin_test?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	// 连接池配置
	SqlDB.SetMaxIdleConns(5)
	SqlDB.SetMaxOpenConns(20)
}
