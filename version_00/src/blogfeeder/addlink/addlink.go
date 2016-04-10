package addlink

import (
	"bufio"
	"comutils"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func AddLinktoAllfiles(dir string, stopic string, topicOK bool, stitle string) {

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	mp := make(map[int]string)

	for i, file := range files {

//		fmt.Println(i, file.Name())
		mp[i] = file.Name()

		//		linkfile := filepath.Join(dir, file.Name())
		//		linktoadd := filepath.Join("/blog/", stopic, stitle+".html")
		//
		//		fileHandle, _ := os.OpenFile(linkfile, os.O_APPEND|os.O_WRONLY, 0666)
		//		writer := bufio.NewWriter(fileHandle)
		//		defer fileHandle.Close()
		//
		//		if topicOK {
		//			linktopicadd := filepath.Join("/blog/", stopic)
		//			fmt.Fprintln(writer, linktopicadd)
		//		}
		//		fmt.Fprintln(writer, linktoadd)
		//
		//		writer.Flush()
	}

	random_file_num := comutils.Random(0, len(mp))

//	fmt.Println(mp[random_file_num])

//	linkfile := filepath.Join(dir, file.Name())
	linkfile := filepath.Join(dir, mp[random_file_num])
	linktoadd := filepath.Join("/blog/", stopic, stitle+".html")

	fileHandle, _ := os.OpenFile(linkfile, os.O_APPEND|os.O_WRONLY, 0666)
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

	if topicOK {
		linktopicadd := filepath.Join("/blog/", stopic)
		fmt.Fprintln(writer, linktopicadd)
	}
	fmt.Fprintln(writer, linktoadd)

	writer.Flush()

}
