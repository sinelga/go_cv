package getLinks

import (
//	"fmt"
//	"github.com/garyburd/redigo/redis"
	"log"
	"log/syslog"
//	"startones"
	"testing"
)

func TestGetAllLinks(t *testing.T) {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	site := "remotejob.work"

//	res := GetAllLinks(golog, c, site)
	GetAllLinks(*golog, site)

//	for _, character := range res {
//
//		fmt.Println(character.Moto)
//
//	}

}
