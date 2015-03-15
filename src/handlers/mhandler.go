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
	//	"fmt"
	//	"log"
//	"log"
	"log/syslog"
	"startones"
//	"toml_parser"
	"sync"
	//	"strings"
)

var startOnce sync.Once

var golog syslog.Writer
var jcv []byte

func MhandleAll(c web.C, w http.ResponseWriter, r *http.Request) {

//	var err error

	golog.Info("UserAgent " + r.UserAgent() + " Host " + r.Host + " RequestURI " + r.RequestURI + " RemoteAddr " + r.RemoteAddr + " referer " + r.Referer()+" Method "+r.Method)

	startOnce.Do(func() {
		golog, jcv = startones.Start()

	})
	

	//	if c.URLParams["id"] != "" {
	//
	//		id := c.URLParams["id"]
	//
	//		for _, item := range bcv.Cv {
	//
	//			if item.Path == id {
	//
	//				if jcv, err = json.Marshal(item); err != nil {
	//
	//					log.Fatal(err.Error())
	//
	//				}
	//
	//			}
	//
	//		}
	//
	//	}

	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding,hashbang")

	w.Write(jcv)

}
