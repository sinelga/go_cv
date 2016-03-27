package addnewblogitem

import (
	"blogfeeder/addlink"
	"domains"
	"encoding/json"
	"io/ioutil"
	"log"
)

func Addnew(blogItems map[string][]domains.BlogItem, item domains.BlogItem, topic string, topicOK bool,linksdir string, filestr string) {

	blogItems[topic] = append(blogItems[topic], item)

	b, err := json.Marshal(blogItems)
	if err != nil {
		log.Println(err)
	}

	ioutil.WriteFile(filestr, b, 0644)

	addlink.AddLinktoAllfiles(linksdir,topic,topicOK, item.Stitle)

}
