package formfeeder

import (
	"domains"
	"encoding/json"
	"fmt"
	"github.com/zenazn/goji/web"
	"net/http"
	"path/filepath"
	//	"strings"
	"blogfeeder/addlink"
	"blogfeeder/addnewblogitem"
	"blogfeeder/check_title"
	"blogfeeder/check_topic"
	"github.com/gosimple/slug"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func HandleForm(c web.C, w http.ResponseWriter, r *http.Request) {
	rootdir := c.Env["rootdir"].(string)
	locale := c.Env["locale"].(string)
	themes := c.Env["themes"].(string)
	//	fmt.Println(rootdir)
	//	fmt.Println(locale)
	//	fmt.Println(themes)

	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/feederform.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		//		fmt.Println("topic:", r.Form["topic"][0])
		//		fmt.Println("title:", r.Form["title"][0])
		//		fmt.Println("contents:", r.Form["contents"][0])

		topic := r.Form["topic"][0]
		title := r.Form["title"][0]
		contents := r.Form["contents"][0]

		if (rootdir != "") && (locale != "") && (themes != "") && (topic != "") && (title != "") && (contents != "") {

			now := time.Now()
			filestr := filepath.Join(rootdir, "dist", locale+"_"+themes+"_"+"blog.json")

			linksdir := filepath.Join(rootdir, "links")

			stitle := slug.Make(title)

			blogItems := make(map[string][]domains.BlogItem)
			item := domains.BlogItem{title, stitle, contents, now, now}

			if _, err := os.Stat(filestr); os.IsNotExist(err) {

				addnewblogitem.Addnew(blogItems, item, topic, true, linksdir, filestr)

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

					stitleOK := check_title.CheckIfExist(stitle, blogItems[topic])
					topicOK := check_topic.CheckTopicExist(topic, blogItems[topic])

					if !stitleOK {

//						fmt.Println("new item", topic, item)
						blogItems[topic] = append(blogItems[topic], item)

						b, err := json.Marshal(blogItems)
						if err != nil {
							log.Println(err)
						}
						ioutil.WriteFile(filestr, b, 0644)

						addlink.AddLinktoAllfiles(linksdir, topic, topicOK, stitle)

					}

				}

			}

		}
		http.Redirect(w, r, "/formfeeder", http.StatusFound)
	}

}
