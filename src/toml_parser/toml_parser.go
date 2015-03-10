package toml_parser

import (
	"domains"
	"fmt"
	 "io/ioutil"
	 "github.com/naoina/toml"
	"os"
)

func Parse(file string) {

	var cv domains.Config
//
//	if _, err := toml.DecodeFile(file, &cv); err != nil {
//		fmt.Println(err)
//		return
//	}
////	fmt.Printf("Title: %s\n", cv.Items)
//	fmt.Println(cv)

	f, err := os.Open(file)
    if err != nil {
        panic(err)
    }
    defer f.Close()
    buf, err := ioutil.ReadAll(f)
    if err != nil {
        panic(err)
    }
//    var config tomlConfig
    if err := toml.Unmarshal(buf, &cv); err != nil {
        panic(err)
    }

	fmt.Println(cv);
	

}
