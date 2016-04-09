package main

import (
	"blogfeeder/addlink"
	"domains"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/SlyMarbo/rss"
	"github.com/gosimple/slug"
	"gopkg.in/gcfg.v1"
	"log"
	"path/filepath"
	"time"
	//	"blogfeeder/addnewblogitem"
	"blogfeeder/check_title"
	"blogfeeder/check_topic"
	"io/ioutil"
	"os"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

var rootdir = ""
var backendrootdir = ""
var locale = ""
var themes = ""

func init() {

	var cfg domains.ServerConfig
	if err := gcfg.ReadFileInto(&cfg, "config.gcfg"); err != nil {
		log.Fatalln(err.Error())

	} else {

		rootdir = cfg.Dirs.Rootdir
		locale = cfg.Main.Locale
		themes = cfg.Main.Themes
		backendrootdir = cfg.Dirs.Backendrootdir

	}

}

func main() {
	flag.Parse() // Scan the arguments list

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)
	}

//	topic := "remote job"
	topic := "javascript"

//	feed, err := rss.Fetch("http://stackoverflow.com/feeds")
	feed, err := rss.Fetch("http://meta.stackexchange.com/feeds/tag?tagnames=javascript&sort=newest")
	if err != nil {
		// handle error.
	}


	items := feed.Items

	for _, item := range items {

		fmt.Println(item.Title)
		title := item.Title
		//		fmt.Println(item.Link)
//		fmt.Println(item.Summary)
		contents := item.Summary

		now := time.Now()
		filestr := filepath.Join(rootdir, "dist", locale+"_"+themes+"_"+"blog.json")

		linksdir := filepath.Join(rootdir, "links")

		stitle := slug.Make(title)
		stopic := slug.Make(topic)

		blogItems := make(map[string][]domains.BlogItem)
		item := domains.BlogItem{stopic, topic, stitle, title, contents, now, now}

		var blogObj map[string]*json.RawMessage
		file, e := ioutil.ReadFile(filestr)
		if e != nil {
			fmt.Printf("File error: %v\n", e)
			os.Exit(1)
		}
		err := json.Unmarshal(file, &blogObj)
		if err != nil {
			fmt.Println("error:", err)
		} else {

			for keytopic, val := range blogObj {

				var items []domains.BlogItem
				err := json.Unmarshal(*val, &items)
				if err != nil {
					fmt.Println("error:", err)
				} else {

					blogItems[keytopic] = items

				}

			}

			key := stopic
			stitleOK := check_title.CheckIfExist(stitle, blogItems[key])
			topicOK := check_topic.CheckTopicExist(topic, blogItems[key])

			if !stitleOK {

				blogItems[key] = append(blogItems[key], item)

				b, err := json.Marshal(blogItems)
				if err != nil {
					log.Println(err)
				}
				ioutil.WriteFile(filestr, b, 0644)

				addlink.AddLinktoAllfiles(linksdir, stopic, topicOK, stitle)

			}

		}
	}

}
