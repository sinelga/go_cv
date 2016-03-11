package dbgetall

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"log"
//	"fmt"
)

func GetAll(db sql.DB,locale string, themes string, table string) []string{


	sqlstr := "select keyword from keywords where locale='" + locale + "' and themes='" + themes + "'"

	rows, err := db.Query(sqlstr)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var outarray []string
	for rows.Next() {
		var keyword string
		rows.Scan(&keyword)
//		fmt.Println(keyword)
		outarray =append(outarray,keyword)
		 
	}
	rows.Close()
	
	return outarray
}
