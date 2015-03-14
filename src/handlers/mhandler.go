package handlers

import (
	"domains"
	"encoding/json"
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
	"log"
	"toml_parser"
	//	"strings"
)

func MhandleAll(c web.C, w http.ResponseWriter, r *http.Request) {

	var jcv []byte
	var err error
	var bcv domains.Config

	//	fmt.Println(r.RequestURI)

	if r.Method == "GET" {

		//	log.Println("method",method)

		bcv = toml_parser.Parse("/home/juno/git/go_cv/cv.toml")

		if jcv, err = json.Marshal(bcv.Cv); err != nil {

			log.Fatal(err.Error())

		}

	}

	if c.URLParams["id"] != "" {

		id := c.URLParams["id"]

		for _, item := range bcv.Cv {

			if item.Path == id {

				if jcv, err = json.Marshal(item); err != nil {

					log.Fatal(err.Error())

				}

			}

		}

	}

	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding,hashbang")

	w.Write(jcv)

}
