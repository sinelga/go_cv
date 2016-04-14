package dbhandler

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"domains"
	"time"
)





func InsertRecord(session mgo.Session, locale string, themes string, site string, menu string) {

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("cv").C("cv")
	
	now :=time.Now()
//	toinsert := []domains.BlogItem{{"111111","1111111","1111111","111111","111111",now,now},{"2222222","2222222","22222","222222","22222",now,now}}
//		
//	toinsertmd :=domains.Md{"en_US","programming","test.com","blog",toinsert}
//	
	
//	err := c.Insert(toinsertmd)
//	if err != nil {
//		log.Fatal(err)
//	}

	result := domains.Md{}
	err := c.Find(bson.M{"locale": "en_US"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result)
	
	newitem :=domains.BlogItem{"44444444444","44444444444444","444444444444","4444444444444","44444444444444",now,now}
	result.Item = append(result.Item,newitem)
	
	err = c.Update(bson.M{"locale": "en_US"},result)
	if err != nil {
		log.Fatal(err)
	}	
	
	
	

}
