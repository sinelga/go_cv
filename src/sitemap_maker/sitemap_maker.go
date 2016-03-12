package main

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"domains"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"log/syslog"
	"mark/dbgetall"
	"math/rand"
	"path/filepath"
	"time"
	//	"startones"
	//	"strconv"
	//	"strings"
	"sitemap_maker/contents_feeder"
	"sitemap_maker/getLinks"
)

var localeFlag = flag.String("locale", "", "must be fi_FI/en_US/it_IT")
var themesFlag = flag.String("themes", "", "must be porno/finance/fortune...")
var contentsdirFlag = flag.String("contentsdir", "", "must dir location contents files")
var linksdirFlag = flag.String("linksdir", "", "must dir location links files")
var mapsdirFlag = flag.String("mapsdir", "", "must dir location sitemaps files")

//var linksdir string
//var  mapsdir string
//var contentsdir  string

//var limitFlag = flag.Int("limit", 0, "if not will be 0")

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
func init() {

}
func main() {
	// Scan the arguments list

	flag.Parse()
	locale := *localeFlag
	themes := *themesFlag
	linksdir := *linksdirFlag
	mapsdir := *mapsdirFlag
	contentsdir := *contentsdirFlag

	if (linksdir != "") && (mapsdir != "") && (contentsdir != "") && (locale != "") && (themes != "") {
		golog, err := syslog.New(syslog.LOG_ERR, "golog")

		defer golog.Close()
		if err != nil {
			log.Fatal("error writing syslog!!")
		}
		//

		db, err := sql.Open("sqlite3", "/home/juno/git/goFastCgi/goFastCgi/singo.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		oldphrases := dbgetall.GetAll(*db, "en_US", "programming", "phrases", "phrase")
		oldkeywords := dbgetall.GetAll(*db, "en_US", "programming", "keywords", "keyword")

		linksmap := getLinks.GetAllLinks(*golog, linksdir)

		docList := new(domains.Pages)
		docList.XmlNS = "http://www.sitemaps.org/schemas/sitemap/0.9"

		var site string

		for key, vals := range linksmap {
			fmt.Println(key)
			fmt.Println(vals)
			site = key

			for _, link := range vals {

				contents_feeder.MakeContents(filepath.Join(contentsdir, site), link,oldkeywords,oldphrases)

				doc := new(domains.Page)
				doc.Loc = "http://" + site + link
				now := time.Now()
				intrand := random(100, 50000)
				minback := time.Duration(intrand)
				lastmod := now.Add(-minback * time.Second)
				doc.Lastmod = lastmod.Format(time.RFC3339)
				doc.Changefreq = "weekly"
				docList.Pages = append(docList.Pages, doc)

			}

			resultXml, err := xml.MarshalIndent(docList, "", "  ")
			if err != nil {

				golog.Crit(err.Error())
			}

			fmt.Println(string(resultXml))
			filestr := mapsdir + "/sitemap_" + site + ".xml"
			ioutil.WriteFile(filestr, resultXml, 0644)
			if err != nil {

				golog.Crit(err.Error())
			}
		}

	} else {
		fmt.Println("try sitemap_maker -h")
	}

}
