package parser

import (
	"fmt"
	"go-spider/engine"
	"go-spider/model"
	"regexp"
)

var itemRe = regexp.MustCompile(
	`<a href="(https://laravelcollections.com/go/[0-9]+)" rel="nofollow" target="_blank">
([^<]*)<small>([^<]*)</small>
</a>
`)

// ParseItem with parse
func ParseItem(contents []byte) []engine.ParseResult {

	matches := itemRe.FindAllSubmatch(contents, -1)
	fmt.Println(len(matches))
	var profileList []engine.ParseResult
	for _, m := range matches {
		profile := model.Profile{}
		profile.Domain = string(m[3])
		profile.Url = string(m[1])
		profile.Title = string(m[2])
		result := engine.ParseResult{}
		result.Items = []interface{}{profile}
		profileList = append(profileList, result)
	}
	return profileList

}
