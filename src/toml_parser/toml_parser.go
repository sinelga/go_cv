package toml_parser

import (
	"domains"
//	"fmt"
	 "io/ioutil"
	 "github.com/naoina/toml"
	"os"
)

func Parse(file string) domains.Config{

	var cv domains.Config

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

//	fmt.Println(cv);
	
	return cv
	

}
