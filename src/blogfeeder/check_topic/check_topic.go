package check_topic

import (
	
//	"fmt"
	"domains"

)

func CheckTopicExist(topic string, items []domains.BlogItem) bool {
	

	if len(items) == 0 {
		
		return true		
		
	}
	
	return false
	
}