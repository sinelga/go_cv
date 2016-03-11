package main

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"encoding/csv"
	"flag"
	"fmt"
	"mark/checkkeyword"
	"mark/dbgetall"
	"os"
	"strings"
	"log"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch

var fileFlag = flag.String("file", "", "file to parse")
var localeFlag = flag.String("locale", "", "must be fi_FI/en_US/it_IT")
var themesFlag = flag.String("themes", "", "must be porno/finance/fortune...")

var mapkeywords map[string]struct{}
var newinsert []string

func main() {
	flag.Parse() // Scan the arguments list

	file := *fileFlag
	locale := *localeFlag
	themes := *themesFlag

	if file != "" {

		db, err := sql.Open("sqlite3", "/home/juno/git/goFastCgiLight/goFastCgiLight/singo.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		mapkeywords = make(map[string]struct{})
		old := dbgetall.GetAll(*db,locale, themes, "keywords")

		for _, val := range old {

			fmt.Println(val)
			mapkeywords[val] = struct{}{}

		}

		csvfile, err := os.Open(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer csvfile.Close()
		reader := csv.NewReader(csvfile)
		reader.LazyQuotes = true

		records, err := reader.ReadAll()
		if err != nil {

			fmt.Println(err)
			return
		} else {

			for _, record := range records {

				keywordsarr := strings.Split(record[0], " ")

				for _, keyword := range keywordsarr {

					if _, ok := mapkeywords[keyword]; ok {

						//						fmt.Println("in map", keyword)

					} else {
						//						fmt.Println("NOT in map", keyword)
						mapkeywords[keyword] = struct{}{}
					}

				}

			}

			for key, _ := range mapkeywords {

				if len(key) > 2 {
					checkkeyword.Check(key)
				}
			}

		}

	} else {
		fmt.Println("try  -h")
	}
}
