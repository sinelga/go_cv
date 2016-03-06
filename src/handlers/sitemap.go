package handlers

import (
	"bytes"
//	"domains"
//	"encoding/xml"
//	"github.com/garyburd/redigo/redis"
//	"handlers/sitemap/createmapfile"
	"io"
//	"log/syslog"
//	"math/rand"
	"net/http"
//	"net/url"
	"os"
//	"sitemap_maker/getLinks"
//	"strconv"
//	"strings"
//	"time"
)


func CheckServeSitemap( w http.ResponseWriter, r *http.Request) {

	site :=  r.Host 

	filestr := "maps/sitemap_" + site + ".xml"

	if _, err := os.Stat(filestr); os.IsNotExist(err) {

		http.NotFound(w, r)

	} else {
		f, _ := os.Open(filestr)
		buf := bytes.NewBuffer(nil)
		io.Copy(buf, f)

		w.Header().Add("Content-type", "application/xml")
		w.Write(buf.Bytes())

	}

}
