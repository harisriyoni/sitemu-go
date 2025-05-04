package app

import (
	"database/sql"

	"github.com/harisriyoni/sitemu-go/helper"
)

func NewDB() *sql.DB {
	dsn := "root@tcp(localhost:3306)/sitemugo"
	db, err := sql.Open("mysql", dsn)
	helper.PanicIfError(err)
	return db
}
