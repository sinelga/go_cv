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
	"log"
	"log/syslog"
	"net/http"
)

var rootdir = ""
var locale = ""
var themes = ""

func init() {

	var cfg domains.ServerConfig
	if err := gcfg.ReadFileInto(&cfg, "config.gcfg"); err != nil {
		log.Fatalln(err.Error())

	} else {

		rootdir = cfg.Dirs.Rootdir
		locale = cfg.Main.Locale
		themes = cfg.Main.Themes

	}

}

func startInit(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		c.Env["rootdir"] = rootdir
		c.Env["locale"] = locale
		c.Env["themes"] = themes		
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)

}

func main() {
	//	goji.Abandon(middleware.Logger)
	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})
	goji.Use(middleware.EnvInit)
	goji.Use(startInit)
	goji.Use(c.Handler)

	goji.Get("/sitemap.xml", handlers.CheckServeSitemap)
	goji.Get("/robots.txt", robots.Generate)
	goji.Get("/formfeeder", formfeeder.HandleForm)
	goji.Post("/formfeeder", formfeeder.HandleForm)	
	goji.Get("/*", handlers.Elaborate)
	//	goji.Get("/echo/json", handlers.Echojson)
	//	goji.Options("/echo/json", handlers.Echojson)
	//	goji.Get("/api", handlers.MhandleAll)
	//	goji.Options("/api", handlers.MhandleAll)
	//	goji.Get("/api/:id", handlers.MhandleAll)
	//	goji.Options("/api/:id", handlers.MhandleAll)

	goji.Serve()

}
