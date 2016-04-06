package handlers

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/zenazn/goji/web"
	"io"
	"net/http"
	"os"
	"strings"
)

type Query struct {
	XMLName xml.Name `xml:"urlset"`
	Locs    []string `xml:"url>loc"`
}
type Loc string

func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func Elaborate(c web.C, w http.ResponseWriter, r *http.Request) {

	rootdir := c.Env["rootdir"].(string)

	exist := false
	sitefull := r.Host
	site := strings.Split(sitefull, ":")[0]
	path := r.URL.Path
	urlstr := "http://" + site + path

	sitemapfile := rootdir + "/maps/sitemap_" + site + ".xml"
	fmt.Println("map", sitemapfile)

	if _, err := os.Stat(sitemapfile); os.IsNotExist(err) {

		//		fmt.Println("NOT EXIST")
		// path/to/whatever does not exist
	} else {

		f, _ := os.Open(sitemapfile)
		buf := bytes.NewBuffer(nil)
		io.Copy(buf, f)

		var q Query
		xml.Unmarshal(buf.Bytes(), &q)

		if stringInSlice(urlstr, q.Locs) {

			exist = true
		}

	}

	if !exist {

		http.NotFound(w, r)

	} else {

		http.ServeFile(w, r, rootdir+"/dist/index.html")
	}

}
