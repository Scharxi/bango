package util

import (
	"database/sql"
	db "go-gin-template/db/sqlc"
)

var Db db.Store

func GetStore() *db.Store {
	return &Db
}

func InitStore(conn *sql.DB) {
	Db = db.NewStore(conn)
}
