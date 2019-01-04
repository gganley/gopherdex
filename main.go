package main

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

func main() {

	db, err := sql.Open("sqlite3", "./poko.db")
	//checkErr(err)
	if err != nil {
		panic(err)
	}
	db.Begin()

	fmt.Println()

}
