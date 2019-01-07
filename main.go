package gopherdex

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

//basic attribute for pokemon
type Pokemon struct {
	Id     int64
	Images map[string]Image
}

//basic attribute for image size
type Image struct {
	URL    string
	Width  int32
	Height int32
}

func Main() {
	db, sqlOpenErr := sql.Open("sqlite3", "./poko.db")
	//checkErr(sqlOpenErr)
	if sqlOpenErr != nil {
		panic(sqlOpenErr)
	}

	tx, dbBeginErr := db.Begin()
	if dbBeginErr != nil {
		panic(dbBeginErr)
	}

	fmt.Println(tx)

}
