package parser

import (
	"fmt"
	"go-spider/engine"
	"go-spider/model"
	"regexp"
)

var itemRe = regexp.MustCompile(`
<li class="link-item list-group-item list-group-item-action">
<h5 class="mb-0">
<a href="(https://laravelcollections.com/go/[0-9]+)" rel="nofollow" target="_blank">([^<]*)<small>([^<]*)</small>
</a>
`)

func ParseItem(contents []byte) engine.ParseResult {
	profile := model.Profile{}
	matches := itemRe.FindAllSubmatch(contents, -1)
	fmt.Println(len(matches))

	result := engine.ParseResult{}
	limit := 2
	for _, m := range matches {
		profile.Domain = string(m[3])
		profile.Url = string(m[1])
		profile.Title = string(m[2])
		limit--
		if limit == 0 {
			break
		}
	}
	result = engine.ParseResult{
		Items: []interface{}{profile},
	}
	fmt.Println(profile)
	return result
}
