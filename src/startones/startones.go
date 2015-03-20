package startones

import (
	"domains"
	"encoding/json"
	"log"
	"log/syslog"
	"toml_parser"
)

var config domains.Config

var bcv domains.Config

var jcv []byte

//func Start(golog syslog.Writer) ([]string,map[string]struct{}) {
func Start() (syslog.Writer, []byte) {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	bcv = toml_parser.Parse("/home/juno/git/go_cv/cv.toml")

	if jcv, err = json.Marshal(bcv.Cv); err != nil {

		log.Fatal(err.Error())

	} else {

		golog.Info("Start Ones")

	}


	return *golog,jcv 

}
