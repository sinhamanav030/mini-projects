package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *sql.DB
	err error
	tpl *template.Template
)

// type UrlMap struct {
// 	Url, Path string
// }

func init() {

	tpl = template.Must(template.ParseFiles("./templates/add.html"))

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

func addInfo(url, path string) {
	query := fmt.Sprintf(`insert into urlmap values("/%s","%s")`, url, path)
	fmt.Println(query)

	stmt, err := db.Prepare(query)

	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := stmt.Exec()
	if err != nil {
		fmt.Println(strings.Split(err.Error(), " "))
		log.Fatal(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(count)
}
