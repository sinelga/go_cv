package domains

import (

)


type Config struct {
    Cv []struct {
        Name string
        Path string
        Item []struct {
            Title string
            Rank int
        }

    }
}





type Cv struct {
	Title string
//	Items []Item 
		
}

type Item struct {
	
	Title string
//	Path string
	Details []Detail
			
}

type Detail struct {
	 Title string
	 Rank int
	
	
}



//type Pages struct {
//	//	Version string   `xml:"version,attr"`
//	XMLName xml.Name `xml:"urlset"`
//	XmlNS   string   `xml:"xmlns,attr"`
//	//	XmlImageNS string   `xml:"xmlns:image,attr"`
//	//	XmlNewsNS  string   `xml:"xmlns:news,attr"`
//	Pages []*Page `xml:"url"`
//}
//
//type Page struct {
//	XMLName    xml.Name `xml:"url"`
//	Loc        string   `xml:"loc"`
//	Lastmod    string   `xml:"lastmod"`
//	Changefreq string   `xml:"changefreq"`
//	//	Name       string   `xml:"news:news>news:publication>news:name"`
//	//	Language   string   `xml:"news:news>news:publication>news:language"`
//	//	Title      string   `xml:"news:news>news:title"`
//	//	Keywords   string   `xml:"news:news>news:keywords"`
//	//	Image      string   `xml:"image:image>image:loc"`
//}
//
//type Config struct {
//	Database struct {
//		ConStr string
//	}
//	Store struct {
//		StoreDir string
//	}
//	Redis struct {
//		Prot string
//		Host string
//	}
//}

//type Chat struct {
//
//	Status string `json: "status"`
//	Known string `json: "known"`
//	Answer string `json: "answer"`
//
//}

//type Character struct {
//	Id               int
//	Name             string
//	Age              int
//	Moto             string
//	Description      string
//	City             string
//	Region_id        int
//	Phone            string
//	Adv_phone_id     int
//	Img_orient       string
//	Topic            string
//	Sex              string
//	Created_at       time.Time
//	Updated_at       time.Time
//	Img_file_name    string
//	Img_content_type string
//	Img_file_size    int
//	Img_updated_at   time.Time
//}
//
//type CharacterRedis struct {
//	Id            int
//	Name          string
//	Age           int
//	Sex           string
//	Moto          string
//	Description   string
//	City          string
//	Region        string
//	Phone         string
//	Created_at    time.Time
//	Img_file_name string
//}
//
//type Paragraph struct {
//	Ptitle     string
//	Pphrase    string
//	Plocallink string
//	Phost      string
//	Sentences  []string
//	Pushsite   string
//}
