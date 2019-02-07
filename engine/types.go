package engine

type Request struct {
	Path       string
	ParserFunc func([]byte) []ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	URL     string
	Type    string
	Id      string
	Payload interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
