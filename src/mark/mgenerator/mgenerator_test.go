package mgenerator

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"mark/dbgetall"
	"testing"
	"log"
)

func TestGenerate(t *testing.T) {

	db, err := sql.Open("sqlite3", "/home/juno/git/goFastCgi/goFastCgi/singo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	oldphrases := dbgetall.GetAll(*db, "en_US", "programming", "phrases", "phrase")
	oldkeywords := dbgetall.GetAll(*db, "en_US", "programming", "keywords", "keyword")	

	Generate(oldkeywords,oldphrases)

}
