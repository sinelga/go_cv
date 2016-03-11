package checkkeyword

import (
	"fmt"
	"regexp"
)

func Check(keyword string) {

	rp := regexp.MustCompile(`[[:punct:]]`)
	rp2 := regexp.MustCompile(`[[:alpha:]]`)
	rp3 := regexp.MustCompile(`[äöøæ]`)

	if rp.MatchString(keyword) {

	} else {
		if rp2.MatchString(keyword) {

			if rp3.MatchString(keyword) {

				fmt.Println("FIN", keyword)

			} else {

//				fmt.Println("Cl", keyword)

			}

		}

	}

}
