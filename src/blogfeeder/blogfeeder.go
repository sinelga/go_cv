package main

import (
	"flag"
	"fmt"
	"os"
//	"domains"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")
var fileFlag = flag.String("file", "", "file to parse")
var localeFlag = flag.String("locale", "", "must be fi_FI/en_US/it_IT")
var themesFlag = flag.String("themes", "", "must be porno/finance/fortune...")
var titleFlag = flag.String("title", "", "must be any string...")

func main() {
	flag.Parse() // Scan the arguments list

	file := *fileFlag
	locale := *localeFlag
	themes := *themesFlag
	title := *titleFlag

	if (file != "") && (locale != "") && (themes != "") && (title != "") {

		if _, err := os.Stat(file); os.IsNotExist(err) {
			fmt.Println(file, "not EXIST")
			// path/to/whatever does not exist
//			blogObj := domains.Blog{"Title","Contents
			
			
		}

	} else {
		fmt.Println("try  -h")
	}

}
