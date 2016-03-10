package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch

var fileFlag = flag.String("file", "", "file to parse")
var mapkeywords map[string][]string

func main() {
	flag.Parse() // Scan the arguments list

	file := *fileFlag

	if file != "" {

		csvfile, err := os.Open(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer csvfile.Close()
		reader := csv.NewReader(csvfile)
		reader.LazyQuotes = true

		records, err := reader.ReadAll()
		if err != nil {

			fmt.Println(err)
			return
		} else {
			//			fmt.Println(records)
			
			mapkeywords=make(map[string][]string)

			for _, record := range records {

//				fmt.Println(record[0])
				keywordsarr := strings.Split(record[0], " ")
//
//				if len(keywordsarr) <= 1 {
//
//					fmt.Println(keywordsarr[0])
////					mapkeywords	["programming"]=append(maplinks[	["programming"],
//
//				} else {

					for _, keyword := range keywordsarr {

//						fmt.Println(keyword)
						mapkeywords	["programming"]=append(mapkeywords["programming"],keyword)						

					}

//				}

				//				fmt.Println(record)
				//				maplinks[site]=append(maplinks[site],record[0])
			}
			
			for _,val :=range mapkeywords {
				
				fmt.Println(val)
			}
			
			//			fmt.Println(maplinks[site])

		}

	} else {
		fmt.Println("try  -h")
	}
}
