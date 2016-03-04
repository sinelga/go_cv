package getLinks

import (
	//	"domains"
	//	"encoding/json"
	//	"github.com/garyburd/redigo/redis"
//	"fmt"
	"log/syslog"
	"os"
	//	"strconv"
	"path/filepath"
	"strings"
)


var links []string

func walkpath(pathstr string, f os.FileInfo, err error) error {

	if f.IsDir() {

		dirscplit := strings.Split(pathstr, "/")

		link := "http:/"

		for i := 6; i < len(dirscplit); i++ {

			link = link + "/" + dirscplit[i]

			if i == len(dirscplit)-1 {
				if i == 6 {

					link = link + "/index.html"
//					fmt.Println(link, i)
					links =append(links,link)
				} else {
					link = link + ".html"
//					fmt.Println(link, i)
					links =append(links,link)

				}

			}

			//			}

		}

		//		fmt.Println(link)

		//		dir, file := path.Split(pathstr)
		//		dirscplit := strings.Split(dir, "/")
		//		id := dirscplit[len(dirscplit)-3]
		//		imgfile := []string{id, file}
		//		imgfiles = append(imgfiles, imgfile)

	}
	return nil
}

func GetAllLinks(golog syslog.Writer, site string) []string {

	rootpath := "/home/juno/git/go_cv/www/" + site

	filepath.Walk(rootpath, walkpath)
	
	return links

	//	limitstr :=strconv.Itoa(limit)

	//	var charactersRedis []domains.CharacterRedis
	//
	//	if bcharactersRedis, err := redis.MultiBulk(c.Do("HVALS", site)); err != nil {
	//
	//		golog.Crit(err.Error())
	//
	//	} else {
	//
	////		fmt.Println(len(bcharactersRedis),site)
	//
	//		for _, bcharacter := range bcharactersRedis {
	//
	//			var v, ok = bcharacter.([]byte)
	//
	//			if ok {
	//
	//				var character domains.CharacterRedis
	//
	//				if err := json.Unmarshal(v, &character); err != nil {
	//					golog.Crit(err.Error())
	//				} else {
	//
	//					if character.Sex == "female" {
	//
	//						charactersRedis = append(charactersRedis, character)
	//
	//					}
	//
	//				}
	//			}
	//
	//		}
	//
	//	}

	//	return charactersRedis
}
