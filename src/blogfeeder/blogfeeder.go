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
	//	"bytes"
	"path/filepath"
	//	"strings"
	"time"
	"blogfeeder/check_title"
	"blogfeeder/addlink"
	"blogfeeder/addnewblogitem"	
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
		filestr := filepath.Join(rootdir,"dist", locale+"_"+themes+"_"+"blog.json")
		
		linksdir := filepath.Join(rootdir,"links")

		stitle := slug.Make(title)

		blogItems := make(map[string][]domains.BlogItem)
		item := domains.BlogItem{title, stitle, "", now, now}

		if _, err := os.Stat(filestr); os.IsNotExist(err) {

			addnewblogitem.Addnew(blogItems,item,topic,linksdir,filestr)


		} else {

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

				stitleOK :=check_title.CheckIfExist(stitle,blogItems[topic])


				if !stitleOK {

					fmt.Println("new item",topic,item)	
					blogItems[topic] = append(blogItems[topic], item)

					b, err := json.Marshal(blogItems)
					if err != nil {
						log.Println(err)
					}
					ioutil.WriteFile(filestr, b, 0644)
										
					addlink.AddLinktoAllfiles(linksdir,topic,stitle)					

				}

			}

		}

	} else {
		fmt.Println("try  -h")
	}

}
