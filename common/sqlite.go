package common

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func InitSqlite() {
	db, err := sql.Open("sqlite3", "./source/data.db")
	if err != nil {
		fmt.Println("hellooooooooooooooo")
		Log("app").Error().Msg(err.Error())
	}
	defer db.Close()

	sql := `create table soup (id integer not null primary key, content text);`
	_, err = db.Exec(sql)
	if err != nil {
		Log("app").Error().Msg(err.Error())
		return
	}

}
