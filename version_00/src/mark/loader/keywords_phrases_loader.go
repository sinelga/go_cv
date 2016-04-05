package main

import (
	_ "github.com/mxk/go-sqlite/sqlite3"
	"database/sql"
	"encoding/csv"
	"flag"
	"fmt"
//	"mark/checkkeyword"
//	"mark/dbgetall"
	"mark/keywords"
	"mark/phrases"	
	"os"
//	"strings"
	"log"
	"domains"
)

var fileFlag = flag.String("file", "", "file to parse")
var localeFlag = flag.String("locale", "", "must be fi_FI/en_US/it_IT")
var themesFlag = flag.String("themes", "", "must be porno/finance/fortune...")

var mapkeywords map[string]struct{}
var newinsert []string
var newobjinsert []domains.KeywordObj

func main() {
	flag.Parse() // Scan the arguments list

	file := *fileFlag
	locale := *localeFlag
	themes := *themesFlag

	if file != "" {

		db, err := sql.Open("sqlite3", "/home/juno/git/goFastCgi/goFastCgi/singo.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()


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
		} 
		
		keywords.Elaborate(locale,themes,*db,records)
		phrases.Elaborate(locale,themes,*db,records)
				

	} else {
		fmt.Println("try  -h")
	}
}
