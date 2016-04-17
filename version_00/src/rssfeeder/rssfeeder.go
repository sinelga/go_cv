package main

import (
	"blogfeeder/addlink"
	"domains"
	"encoding/csv"
	//	"encoding/json"
	//	"flag"
	"fmt"
	"github.com/SlyMarbo/rss"
	"github.com/gosimple/slug"
	"gopkg.in/gcfg.v1"
	"log"
	"path/filepath"
	//	"time"
	//	"blogfeeder/addnewblogitem"
	//	"blogfeeder/check_title"
	//	"blogfeeder/check_topic"
	//	"io/ioutil"
	"dbhandler"
	"gopkg.in/mgo.v2"
	"os"
	"time"
)

var rootdir = ""
var backendrootdir = ""
var locale = ""
var themes = ""
var rssresorsesfile = ""

//var session *mgo.Session

var resorses []domains.Rssresors

func init() {

	var cfg domains.ServerConfig
	if err := gcfg.ReadFileInto(&cfg, "config.gcfg"); err != nil {
		log.Fatalln(err.Error())

	} else {

		rootdir = cfg.Dirs.Rootdir
		locale = cfg.Main.Locale
		themes = cfg.Main.Themes
		backendrootdir = cfg.Dirs.Backendrootdir
		rssresorsesfile = cfg.Dirs.Rssresorsesfile

	}

	csvfile, err := os.Open(rssresorsesfile)
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

	for _, record := range records {

		//		fmt.Println(record[0],record[1])
		res := domains.Rssresors{record[0], record[1]}
		resorses = append(resorses, res)
	}

}

func main() {

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	linksdir := filepath.Join(rootdir, "links")

	uniqstitle := dbhandler.GetAllStitle(*session, locale, themes)

	for _, res := range resorses {

		now := time.Now()

		topic := res.Topic
		stopic := slug.Make(topic)
		fmt.Println(topic)

		feed, err := rss.Fetch(res.Link)
		if err != nil {
			// handle error.
			panic(err.Error())
		}

		items := feed.Items

		for _, item := range items {

			title := item.Title
			stitle := slug.Make(title)

			if _, ok := uniqstitle[stitle]; !ok {

				fmt.Println("new stitle", stitle)
				contents := item.Summary

				site := addlink.AddLinktoAllfiles(linksdir, stopic, stitle)
				
				fmt.Println("site",site)
				item := domains.BlogItem{stopic, topic, stitle, title, contents, now, now}

				dbhandler.InsertRecord(*session, locale, themes, site, "blog", stopic,topic, item)

			}

			//			now := time.Now()
			//			filestr := filepath.Join(rootdir, "dist", locale+"_"+themes+"_"+"blog.json")
			//
			//			linksdir := filepath.Join(rootdir, "links")
			//
			//			stitle := slug.Make(title)
			//			stopic := slug.Make(topic)
			//
			//			blogItems := make(map[string][]domains.BlogItem)
			//			item := domains.BlogItem{stopic, topic, stitle, title, contents, now, now}
			//
			//			var blogObj map[string]*json.RawMessage
			//			file, e := ioutil.ReadFile(filestr)
			//			if e != nil {
			//				fmt.Printf("File error: %v\n", e)
			//				os.Exit(1)
			//			}
			//			err := json.Unmarshal(file, &blogObj)
			//			if err != nil {
			//				fmt.Println("error:", err)
			//			} else {
			//
			//				for keytopic, val := range blogObj {
			//
			//					var items []domains.BlogItem
			//					err := json.Unmarshal(*val, &items)
			//					if err != nil {
			//						fmt.Println("error:", err)
			//					} else {
			//
			//						blogItems[keytopic] = items
			//
			//					}
			//
			//				}
			//
			//				key := stopic
			//				stitleOK := check_title.CheckIfExist(stitle, blogItems[key])
			//				topicOK := check_topic.CheckTopicExist(topic, blogItems[key])
			//
			//				if !stitleOK {
			//
			//					blogItems[key] = append(blogItems[key], item)
			//
			//					b, err := json.Marshal(blogItems)
			//					if err != nil {
			//						log.Println(err)
			//					}
			//					ioutil.WriteFile(filestr, b, 0644)
			//
			//					addlink.AddLinktoAllfiles(linksdir, stopic, topicOK, stitle)
			//
			//				}
			//
			//			}
		}
	}

}
