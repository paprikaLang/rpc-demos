package engine

import (
	"go-spider/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetching %s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher ; error "+"fetching url %s: %v", r.Url, err)
			continue
		}
		parseRes := r.ParserFunc(body)
		requests = append(requests, parseRes.Requests...)

		for _, item := range parseRes.Items {
			log.Printf("got item %v", item)
		}

	}
}
