package parser

import (
	"go-spider/engine"
	"regexp"
)

var topicRe = regexp.MustCompile(`
<div class="sidebar-item">
<a href="(https://laravelcollections.com/[^"]*)">([^<]*)</a>
<span class="float-right badge badge-secondary">([0-9]+)</span>
</div>`)

// ParseTopic with parse
func ParseTopic(contents []byte) []engine.ParseResult {
	matches := topicRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	var topicList []engine.ParseResult
	for _, m := range matches {
		// result.Items = append(result.Items, "Topic "+string(m[2]))
		// result.Items = append(result.Items, "Total "+string(m[3]))
		topic := string(m[2])
		result.Requests = append(result.Requests, engine.Request{
			Path:   string(m[1]),
			Parser: NewItemParser(topic),
		})
	}
	return append(topicList, result)
}

// func ItemParser(topic string) engine.ParserFunc {
// 	return func(contents []byte) []engine.ParseResult {
// 		return ParseItem(contents, topic)
// 	}
// }
