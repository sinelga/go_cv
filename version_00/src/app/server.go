package main

import (
	"domains"
	"formfeeder"
	"github.com/rs/cors"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
	"gopkg.in/gcfg.v1"
	"handlers"
	"handlers/robots"
	"handlers/blog"
	"log"
//	"log/syslog"
	"net/http"
	"flag"
)

var rootdir = ""
var rootdirm = ""
var backendrootdir = ""
var locale = ""
var themes = ""

func init() {

	var cfg domains.ServerConfig
	if err := gcfg.ReadFileInto(&cfg, "config.gcfg"); err != nil {
		log.Fatalln(err.Error())

	} else {

		rootdir = cfg.Dirs.Rootdir
		rootdirm = cfg.Dirs.Rootdirm
		locale = cfg.Main.Locale
		themes = cfg.Main.Themes
		backendrootdir = cfg.Dirs.Backendrootdir

	}

}

func startInit(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		c.Env["rootdir"] = rootdir
		c.Env["rootdirm"] = rootdirm		
		c.Env["locale"] = locale
		c.Env["themes"] = themes
		c.Env["backendrootdir"] = backendrootdir
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)

}

func main() {
	//	goji.Abandon(middleware.Logger)
//	golog, err := syslog.New(syslog.LOG_ERR, "golog")
//
//	defer golog.Close()
//	if err != nil {
//		log.Fatal("error writing syslog!!")
//	}

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},		
	})
	goji.Use(middleware.EnvInit)
	goji.Use(startInit)
	goji.Use(c.Handler)

	goji.Get("/sitemap.xml", handlers.CheckServeSitemap)
	goji.Get("/robots.txt", robots.Generate)
	goji.Get("/formfeeder", formfeeder.HandleForm)
	goji.Post("/formfeeder", formfeeder.HandleForm)
	goji.Get("/api/blog/:stopic/:stitle",blog.GetIemDetails)	
	goji.Get("/api/blog/:stopic",blog.GetItem)
	goji.Get("/api/blog",blog.BlogIndex)
	goji.Get("/*", handlers.Elaborate)

    flag.Set("bind", ":8001")
	goji.Serve()

}
