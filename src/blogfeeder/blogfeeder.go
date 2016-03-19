package main

import (
	"domains"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gosimple/slug"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

var rootdirFlag = flag.String("rootdir", "", "file to parse")
var localeFlag = flag.String("locale", "", "must be fi_FI/en_US/it_IT")
var themesFlag = flag.String("themes", "", "must be porno/finance/fortune...")
var topicFlag = flag.String("topic", "", "must string like programming, java, c ..")
var titleFlag = flag.String("title", "", "must be any string...")

func main() {
	flag.Parse() // Scan the arguments list

	rootdir := *rootdirFlag
	locale := *localeFlag
	themes := *themesFlag
	topic := *topicFlag
	title := *titleFlag

	if (rootdir != "") && (locale != "") && (themes != "") && (topic != "") && (title != "") {

		now := time.Now()
		file := filepath.Join(rootdir, "blog.json")

		fmt.Println(file)
		stitle := slug.Make(title)

		jsonfiledir := filepath.Join(rootdir, topic, stitle)
		fmt.Println(jsonfiledir)

		if _, err := os.Stat(file); os.IsNotExist(err) {
			fmt.Println(file, "not EXIST")

			blogItems := make(map[string][]domains.BlogItem)
			item := domains.BlogItem{stitle, "", now, now}
			blogItems[topic] = append(blogItems[topic], item)

			b, err := json.Marshal(blogItems)
			if err != nil {
				log.Println(err)
			}

			ioutil.WriteFile(file, b, 0644)

		} else {

			//			var blogObj  domains.Blog
			var blogObj map[string]*json.RawMessage
			//			var blogObj  map[string]*[]struct{}
			file, e := ioutil.ReadFile(file)
			if e != nil {
				fmt.Printf("File error: %v\n", e)
				os.Exit(1)
			}
			err := json.Unmarshal(file, &blogObj)
			if err != nil {
				fmt.Println("error:", err)
			} else {

				fmt.Println(blogObj)
				for key, val := range blogObj {
					fmt.Println(key, val)
					
//					var 
					var items []domains.BlogItem
					
					
//
					err := json.Unmarshal(*val, &items)
					if err != nil {
						fmt.Println("error:", err)
					} else {
						fmt.Println(items)
						
					}

				}

			}

		}

		// path/to/whatever does not exist
		//			blogObj := domains.Blog{"Title","Contents

	} else {
		fmt.Println("try  -h")
	}

}
