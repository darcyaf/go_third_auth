package tools

import (
	"fmt"
	"net/url"
)

func MapMerge(maps ...map[string]interface{}) url.Values {
	var values = url.Values{}
	for _, m := range maps {
		for k, v := range m {
			values.Add(k, fmt.Sprint(v))
		}
	}
	return values
}