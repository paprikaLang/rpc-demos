package engine

import (
	"log"
)

type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		for _, perResult := range parseResult {
			requests = append(requests, perResult.Requests...)

			for _, item := range perResult.Items {
				log.Printf("got item %v", item)
			}
		}

	}
}
