package keywords

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"domains"
	//	"fmt"
	"mark/dbgetall"
	"mark/keywords/checkkeyword"
	"mark/keywords/insertnewkeywords"
	"strings"
)

var mapkeywords map[string]struct{}
var newinsert []string
var newobjinsert []domains.KeywordObj

func Elaborate(locale string, themes string, db sql.DB, records [][]string) {

	mapkeywords = make(map[string]struct{})
	old := dbgetall.GetAll(db, locale, themes, "keywords")

	for _, val := range old {

		mapkeywords[val] = struct{}{}

	}

	for _, record := range records {
		keywordsarr := strings.Split(record[0], " ")

		for _, keyword := range keywordsarr {

			if _, ok := mapkeywords[keyword]; ok {

			} else {
				//				fmt.Println("NOT in map", keyword)
				mapkeywords[keyword] = struct{}{}
				newinsert = append(newinsert, keyword)
			}

		}

	}

	for _, insertstr := range newinsert {

		if len(insertstr) > 2 {
			keywordObj := checkkeyword.Check(insertstr)
			if keywordObj.Keyword != "" {
				newobjinsert = append(newobjinsert, keywordObj)

			}

		}
	}
	insertnewkeywords.InsertAll(db, locale, themes, newobjinsert)

}
