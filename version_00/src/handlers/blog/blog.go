package blog

import (
	"domains"
	"encoding/json"
	"fmt"
	"github.com/zenazn/goji/web"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strings"
)



func BlogIndex(w http.ResponseWriter, r *http.Request) {

	locale := "en_US"
	themes := "programming"
	site := "127.0.0.1"
	menu := "blog"

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("cv").C("cv")

	var results []domains.Md

	query := bson.M{"locale": locale, "themes": themes, "site": site, "menu": menu}

	err = c.Find(query).All(&results)

	var blogindex []domains.BlogIndex
	//	fmt.Println(results)
	for _, result := range results {

		fmt.Println(result.Topic)
		indexItem := domains.BlogIndex{result.Stopic, result.Topic}

		blogindex = append(blogindex, indexItem)

	}
	encoder := json.NewEncoder(w)
	encoder.Encode(blogindex)
}

func GetItem(c web.C, w http.ResponseWriter, r *http.Request) {

	locale := "en_US"
	themes := "programming"
	site := "127.0.0.1"
	menu := "blog"
	stopic := c.URLParams["stopic"]
	//	fmt.Println(c.URLParams["stopic"])

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	cm := session.DB("cv").C("cv")

	var results []domains.Md


	query := bson.M{"locale": locale, "themes": themes, "site": site, "menu": menu, "stopic": stopic}

	err = cm.Find(query).All(&results)

	encoder := json.NewEncoder(w)
	encoder.Encode(results)

}

func GetIemDetails(c web.C, w http.ResponseWriter, r *http.Request) {
	locale := "en_US"
	themes := "programming"
	site := "127.0.0.1"
	menu := "blog"
	stopic := c.URLParams["stopic"]
	stitle := c.URLParams["stitle"]
	fmt.Println(c.URLParams["stitle"])

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	cm := session.DB("cv").C("cv")

	var results []domains.Md

//db.cv.find({locale : "en_US", themes : "programming", site: "127.0.0.1",menu: "blog",stopic: "remote-job"},{items: { $elemMatch: {stitle:  "do-c-compilers-and-nvcc-respect-the-__restrict__-keyword-within-structs"}}})

	query := bson.M{"locale": locale, "themes": themes, "site": site, "menu": menu, "stopic": stopic}
	err = cm.Find(query).All(&results)
	if err != nil {
		fmt.Println(err.Error())
		
	}

	var selecteditem domains.BlogItem
	
	for _,item :=range results[0].Items {
		if strings.HasPrefix(item.Stitle,stitle) {
			
			selecteditem = domains.BlogItem{item.Stopic,item.Topic,item.Stitle,item.Title,item.Contents,item.Created_at,item.Updated_at}

		}
		
		
	}
	
	
	
	encoder := json.NewEncoder(w)
	encoder.Encode(selecteditem)

}
