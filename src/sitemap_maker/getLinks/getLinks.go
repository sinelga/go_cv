package getLinks

import (
//	"domains"
//	"encoding/json"
//	"github.com/garyburd/redigo/redis"
	"log/syslog"
	"os"	
//	"fmt"
	//	"strconv"
)

func walkpath(pathstr string, f os.FileInfo, err error) error {

	if !f.IsDir() {

//		dir, file := path.Split(pathstr)
//		dirscplit := strings.Split(dir, "/")
//		id := dirscplit[len(dirscplit)-3]
//		imgfile := []string{id, file}
//		imgfiles = append(imgfiles, imgfile)

	}
	return nil
}


func GetAllLinks(golog syslog.Writer, site string) {

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
