package addlink

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func AddLinktoAllfiles(dir string, topic string, stitle string) {

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		linkfile := filepath.Join(dir, file.Name())
		linktoadd := filepath.Join("/blog/", topic, stitle+".html")
		fmt.Println(linktoadd)

		fileHandle, _ := os.OpenFile(linkfile, os.O_APPEND|os.O_WRONLY, 0666)
		writer := bufio.NewWriter(fileHandle)
		defer fileHandle.Close()

		fmt.Fprintln(writer, linktoadd)
		writer.Flush()
	}

}
