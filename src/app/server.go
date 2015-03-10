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

	goji.Get("/api", handlers.MhandleAll)
	
	goji.Serve()
	

}

