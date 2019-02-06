package parser

import (
	"go-spider/engine"
	"regexp"
)

var topicRe = regexp.MustCompile(`
<div class="sidebar-item">
<a href="https://laravelcollections.com/([^>]*)">([^<]*)</a>
<span class="float-right badge badge-secondary">([0-9]+)</span>
</div>`)

func ParseTopic(contents []byte) engine.ParseResult {
	matches := topicRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	limit := 3
	for _, m := range matches {
		result.Items = append(result.Items, "Topic "+string(m[2]))
		result.Items = append(result.Items, "Total "+string(m[3]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        "https://laravelcollections.com/" + string(m[1]),
			ParserFunc: ParseItem,
		})
		limit--
		if limit == 0 {
			break
		}
	}
	return result
}
