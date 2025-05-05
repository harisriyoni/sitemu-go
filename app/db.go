package app

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/harisriyoni/sitemu-go/helper"
)

func NewDB() *sql.DB {
	// Format DSN: username:password@tcp(host:port)/dbname?options
	dsn := "sitemugo_amloudcall:0fe405b996bc33e862b98c2c2c6a3d1aacb29002@tcp(w1mdb.h.filess.io:3307)/sitemugo_amloudcall?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", dsn)
	helper.PanicIfError(err)

	err = db.Ping() // Check connection
	helper.PanicIfError(err)

	return db
}
