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
	
	return cv	

}

func ParseWorkPlaces(file string) domains.Job{

	var job domains.Job

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
    if err := toml.Unmarshal(buf, &job); err != nil {
        panic(err)
    }
	
	return job	

}
