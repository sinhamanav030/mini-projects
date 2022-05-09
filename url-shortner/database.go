package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	db  *sql.DB
	err error
)

// type UrlMap struct {
// 	Url, Path string
// }

func init() {
	db, err = sql.Open("mysql", "root:my-secret-pw@tcp(localhost:3306)/urlShortner?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connection Established")

}

func getInfo() map[string]string {
	rows, err := db.Query(`select * from urlmap`)
	if err != nil {
		log.Fatal(err)
	}

	urlmap := make(map[string]string)

	for rows.Next() {
		var url, path string
		err = rows.Scan(&url, &path)
		if err != nil {
			log.Fatal(err)
		}
		urlmap[url] = path

	}
	return urlmap
}
