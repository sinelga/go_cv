package handlers

import (
//	"domains"
//	"encoding/json"
	"github.com/zenazn/goji/web"
//	"github.com/garyburd/redigo/redis"
//	"handlers/getAll"
//	"handlers/getOne"
//	"log/syslog"
	"net/http"
//	"startones"
//	"sync"
//	"net/url"
//	"net"
//	"strings"
)

//var startOnce sync.Once
//var golog syslog.Writer
//
//var config domains.Config

func MhandleAll(c web.C, w http.ResponseWriter, r *http.Request) {

//	startOnce.Do(func() {
//		golog, config = startones.Start()
//
//	})
//	
//	site, _, _ := net.SplitHostPort(r.Host)
//			
//	if site == "localhost" {
//		
//		site="www.test.com"
//	}
//	
//	if strings.HasPrefix(site,"192.168.") {
//		
//		site="www.test.com"
//	}	
//	
//	golog.Info("Site "+site)
//	
//	rds, err := redis.Dial("tcp", ":6379")
//	if err != nil {
//
//		golog.Crit(err.Error())
//
//	}
//	defer rds.Close()
//
//	id := c.URLParams["id"]
//
//	var bytes []byte
//	var e error
//
//	if id == "" {
//
//		characters := getAll.GetAll(golog,rds,site)
//
//		bytes, e = json.Marshal(characters)
//		if e != nil {
//
//			golog.Err(e.Error())
//
//		}
//
//	} else {
//
//
//		character := getOne.GetById(golog,rds,site, id)
//
//		bytes, e = json.Marshal(character)
//		if e != nil {
//
//			golog.Err(e.Error())
//
//		}
//
//	}
//
//	w.Write(bytes)

}
