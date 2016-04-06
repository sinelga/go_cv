package main

import (
	"database/sql"
	"domains"
	"encoding/xml"
	"flag"
	"fmt"
	_ "github.com/mxk/go-sqlite/sqlite3"
	"io/ioutil"
	"log"
	"log/syslog"
	"mark/dbgetall"
	"math/rand"
	"path/filepath"
	"sitemap_maker/contents_feeder"
	"sitemap_maker/getLinks"
	"sitemap_maker/unmarshalsitemap"
	"time"
)

var localeFlag = flag.String("locale", "", "must be fi_FI/en_US/it_IT")
var themesFlag = flag.String("themes", "", "must be porno/finance/fortune...")
var dblocFlag = flag.String("dbloc", "", "must be somthing like en_US_programming.db")
var contentsdirFlag = flag.String("contentsdir", "", "must dir location contents files")
var linksdirFlag = flag.String("linksdir", "", "must dir location links files")
var mapsdirFlag = flag.String("mapsdir", "", "must dir location sitemaps files")

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
	dbloc := *dblocFlag

	if (linksdir != "") && (mapsdir != "") && (contentsdir != "") && (locale != "") && (themes != "") && (dbloc != "") {
		golog, err := syslog.New(syslog.LOG_ERR, "golog")

		defer golog.Close()
		if err != nil {
			log.Fatal("error writing syslog!!")
		}
		//

		db, err := sql.Open("sqlite3", dbloc)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		oldphrases := dbgetall.GetAll(*db, locale, themes, "phrases", "phrase")
		oldkeywords := dbgetall.GetAll(*db, locale, themes, "keywords", "keyword")

		linksmap := getLinks.GetAllLinks(*golog, linksdir)

		docList := new(domains.Pages)
		docList.XmlNS = "http://www.sitemaps.org/schemas/sitemap/0.9"

		var site string

		for key, vals := range linksmap {

			site = key
			filestr := mapsdir + "/sitemap_" + site + ".xml"

			fmt.Println("Site", site, filestr)

			sitemapObjs := unmarshalsitemap.Get(filestr)

			stitlemap := make(map[string]struct{})

			var oldlinks []string
			var durationfixed float64
			durationfixed = float64(0)

			for _, sitemapObj := range sitemapObjs {

				stitlemap[sitemapObj.Loc] = struct{}{}

				if sitemapObj.Changefreq == "monthly" {

					durationfixed = float64(21)

				}
				//			    fmt.Println(sitemapObj.Changefreq)

				if sitemapObj.Hoursduration > durationfixed {

					fmt.Println("need update ", sitemapObj.Loc)
					oldlinks = append(oldlinks, sitemapObj.Loc)
				}

			}

			for _, link := range vals {

				pageurl := "http://" + site + link

				if _, ok := stitlemap[pageurl]; !ok {

					fmt.Println("new link", pageurl)

				}

				contents_feeder.MakeContents(filepath.Join(contentsdir, site), link, oldkeywords, oldphrases)

				doc := new(domains.Page)
				doc.Loc = pageurl
				now := time.Now()
				intrand := random(100, 50000)
				minback := time.Duration(intrand)
				lastmod := now.Add(-minback * time.Second)
				doc.Lastmod = lastmod.Format(time.RFC3339)
				doc.Changefreq = "monthly"
				docList.Pages = append(docList.Pages, doc)

			}

			resultXml, err := xml.MarshalIndent(docList, "", "  ")
			if err != nil {

				golog.Crit(err.Error())
			}
			docList.Pages = nil

			ioutil.WriteFile(filestr, resultXml, 0644)
			if err != nil {

				golog.Crit(err.Error())
			}
		}

	} else {
		fmt.Println("try sitemap_maker -h")
	}

}
