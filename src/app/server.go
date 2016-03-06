package main 

import (
	"log"
	"log/syslog"
	"github.com/rs/cors"
	"github.com/zenazn/goji"
	"handlers"
)

func main() {
	
	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}
	
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, 
	})
	goji.Use(c.Handler)

	goji.Get("/sitemap.xml",handlers.CheckServeSitemap)
	goji.Get("/*",handlers.Elaborate) 
//	goji.Get("/echo/json", handlers.Echojson)
//	goji.Options("/echo/json", handlers.Echojson)
//	goji.Get("/api", handlers.MhandleAll)
//	goji.Options("/api", handlers.MhandleAll)
//	goji.Get("/api/:id", handlers.MhandleAll)
//	goji.Options("/api/:id", handlers.MhandleAll)	
			
	goji.Serve()
	

}

