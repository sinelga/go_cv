package main

import (
	"domains"
	"flag"
	"fmt"
	"log"
	"log/syslog"
	"math/rand"
	"io/ioutil"
	"time"
	//	"fmt"
	"encoding/xml"
	//	"startones"
	//	"strconv"
	//	"strings"

	"sitemap_maker/getLinks"
)

var siteFlag = flag.String("site", "", "must be test.com www.test.com")

//var limitFlag = flag.Int("limit", 0, "if not will be 0")

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func main() {
	flag.Parse() // Scan the arguments list

	site := *siteFlag

	if site != "" {
		golog, err := syslog.New(syslog.LOG_ERR, "golog")

		defer golog.Close()
		if err != nil {
			log.Fatal("error writing syslog!!")
		}

		//	golog, _ := startones.Start()

		links := getLinks.GetAllLinks(*golog, site)

		docList := new(domains.Pages)
		docList.XmlNS = "http://www.sitemaps.org/schemas/sitemap/0.9"

		for _, link := range links {

			doc := new(domains.Page)
			doc.Loc = link
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
		
		filestr := "/home/juno/git/go_cv/maps/sitemap_"+site+".xml"
		
		ioutil.WriteFile(filestr,resultXml, 0644)
		if err != nil {

			golog.Crit(err.Error())
		}				

		//
		//	var Url *url.URL
		//
		//	docList := new(domains.Pages)
		//	docList.XmlNS = "http://www.sitemaps.org/schemas/sitemap/0.9"
		//
		//	for i := 0; i < limit; i++ {
		//
		//		Url, err = url.Parse("http://" + site)
		//		if err != nil {
		//			golog.Crit(err.Error())
		//		}
		//
		//		permlink :=strings.Split(characters[i].Moto," ")
		////		Url.Path += "/"+ strconv.Itoa(characters[i].Id) + "/" +permlink[0]+"-"+permlink[1]+".html"
		//		Url.Path += "/"+ characters[i].Id + "/" +permlink[0]+"-"+permlink[1]+".html"
		//
		//		doc := new(domains.Page)
		//		doc.Loc = Url.String()
		//		doc.Lastmod =characters[i].Created_at.Format(time.RFC3339)
		//		doc.Changefreq = "weekly"
		//
		//		docList.Pages = append(docList.Pages, doc)
		//
		//	}
		//
		//	resultXml, err := xml.MarshalIndent(docList, "", "  ")
		//	if err != nil {
		//
		//		golog.Crit(err.Error())
		//	}
		//
		//	fmt.Println(string(resultXml))

	} else {
		fmt.Println("try sitemap_maker -h")
	}

}
