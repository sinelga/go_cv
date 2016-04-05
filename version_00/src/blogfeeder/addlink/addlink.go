package addlink

import (
	"bufio"
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

	for _, file := range files {

		linkfile := filepath.Join(dir, file.Name())
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

}
