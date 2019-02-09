package engine

import (
	"go-spider/fetcher"
	"log"
)

func Worker(r Request) ([]ParseResult, error) {
	log.Printf("Fetching %s", r.Path)
	body, err := fetcher.Fetch(r.Path)
	if err != nil {
		log.Printf("Fetcher ; error "+"fetching url %s: %v", r.Path, err)
		return []ParseResult{}, err
	}

	return r.Parser.Parse(body), nil
}
