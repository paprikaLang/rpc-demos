package parser

import (
	"fmt"
	"go-spider/engine"
	"go-spider/model"
	"regexp"
)

var itemRe = regexp.MustCompile(
	`<a href="https://laravelcollections.com/go/([0-9]+)" rel="nofollow" target="_blank">
([^<]*)<small>([^<]*)</small>
</a>
`)

// ParseItem with parse
func ParseItem(contents []byte, topic string) []engine.ParseResult {

	matches := itemRe.FindAllSubmatch(contents, -1)
	fmt.Println(len(matches))
	var profileList []engine.ParseResult
	for _, m := range matches {
		profile := model.Profile{}
		profile.Domain = string(m[3])
		profile.Topic = topic
		profile.Title = string(m[2])
		result := engine.ParseResult{}
		result.Items = []engine.Item{
			{
				URL:     "https://laravelcollections.com/go/" + string(m[1]),
				Id:      string(m[1]),
				Type:    "collections",
				Payload: profile,
			},
		}
		profileList = append(profileList, result)
	}
	return profileList

}

type ItemParser struct {
	topic string
}

func (f *ItemParser) Parse(contents []byte) []engine.ParseResult {
	return ParseItem(contents, f.topic)
}

func (f *ItemParser) Serialize() (name string, args interface{}) {
	return "ItemParser", f.topic
}
func NewItemParser(topic string) *ItemParser {
	return &ItemParser{
		topic: topic,
	}
}
