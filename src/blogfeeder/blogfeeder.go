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
	"time"

	//	"encoding/binary"
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
		filestr := filepath.Join(rootdir, locale+"_"+themes+"_"+"blog.json")

		stitle := slug.Make(title)

		blogItems := make(map[string][]domains.BlogItem)
		item := domains.BlogItem{stitle, "", now, now}

		if _, err := os.Stat(filestr); os.IsNotExist(err) {

			blogItems[topic] = append(blogItems[topic], item)

			b, err := json.Marshal(blogItems)
			if err != nil {
				log.Println(err)
			}

			ioutil.WriteFile(filestr, b, 0644)

		} else {

			var blogObj map[string]*json.RawMessage
			//			var blogObj map[string]*[]struct{}
			file, e := ioutil.ReadFile(filestr)
			if e != nil {
				fmt.Printf("File error: %v\n", e)
				os.Exit(1)
			}
			err := json.Unmarshal(file, &blogObj)
			if err != nil {
				fmt.Println("error:", err)
			} else {

//				fmt.Println(blogObj)
				for key, val := range blogObj {
//					fmt.Println(key, val)
					var items []domains.BlogItem
					err := json.Unmarshal(*val, &items)
					if err != nil {
						fmt.Println("error:", err)
					} else {
//						fmt.Println(items)

						blogItems[key] = items

					}

				}

				
//				if val, ok := blogItems[topic]; ok {
//    				fmt.Println("topic  Exist")
//    				newval := append(val, item)    				
//    				blogItems[topic] = newval
//    				
//				} else {
					
					blogItems[topic] =append(blogItems[topic],item)					
					
//				}

				b, err := json.Marshal(blogItems)
				if err != nil {
					log.Println(err)
				}
//				fmt.Println(string(b))
				ioutil.WriteFile(filestr, b, 0644)

			}

		}

	} else {
		fmt.Println("try  -h")
	}

}
