package contents_feeder

import (
	"domains"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mark/mgenerator"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func MakeContents(contentsdir string, pathstr string, keywords []string, phrases []string) {

	//	fmt.Println(rootdir, pathstr)

	splitpath := strings.Split(pathstr, "/")

	directory := contentsdir

	for i := 0; i < len(splitpath)-1; i++ {

		//		directory = directory + "/" + splitpath[i]
		directory = filepath.Join(directory, splitpath[i])

	}
	lastspilt := splitpath[len(splitpath)-1]
	//	directory = directory + "/" + strings.Split(lastspilt, ".")[0]
	directory = filepath.Join(directory, strings.Split(lastspilt, ".")[0])
	//	fmt.Println(directory)
	jsonpath := directory + "/" + lastspilt + ".json"

	if _, err := os.Stat(directory); os.IsNotExist(err) {

		fmt.Println(directory, "NOT exist")
		if err := os.MkdirAll(directory, 0777); err != nil {
			fmt.Println(err.Error())
		} else {

			//			contents := domains.Contents{"title", "moto", "contents", time.Now(), time.Now()}
			contents := mgenerator.Generate(keywords, phrases)

			b, err := json.Marshal(contents)
			if err != nil {
				log.Println(err)
			}

			ioutil.WriteFile(jsonpath, b, 0644)

		}
		// path/to/whatever does not exist
	} else {

		fmt.Println(directory, "EXist")

		if _, err := os.Stat(jsonpath); os.IsNotExist(err) {

			contents := domains.Contents{"title", "moto", "contents", time.Now(), time.Now()}

			b, err := json.Marshal(contents)
			if err != nil {
				log.Println(err)
			}

			ioutil.WriteFile(jsonpath, b, 0644)

		} else {

			byt, err := ioutil.ReadFile(jsonpath)
			if err != nil {
				fmt.Print("Error:", err)
			}

			var dat domains.Contents
			if err := json.Unmarshal(byt, &dat); err != nil {
				panic(err)
			}
//			fmt.Println(dat)
			contents := domains.Contents{dat.Title, dat.Moto, dat.Contents, dat.Created_at, time.Now()}

			b, err := json.Marshal(contents)
			if err != nil {
				log.Println(err)
			}

			ioutil.WriteFile(jsonpath, b, 0644)

		}
	}

	//	fulldir := rootdir+strings.Split(pathstr,"/")[0]
	//
	//	fmt.Println(fulldir	)

}
