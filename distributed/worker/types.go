package worker

import (
	"errors"
	"fmt"
	"go-spider/engine"
	"go-spider/main/parser"
	"log"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

type Request struct {
	URL    string
	Parser SerializedParser
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		URL: r.Path,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(result []engine.ParseResult) []ParseResult {
	var array []ParseResult
	for _, r := range result {
		result := ParseResult{
			Items: r.Items,
		}
		for _, req := range r.Requests {
			result.Requests = append(result.Requests, SerializeRequest(req))
		}
		array = append(array, result)
	}
	return array
}

func DeserializeRequest(r Request) (engine.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, nil
	}
	return engine.Request{
		Path:   r.URL,
		Parser: parser,
	}, nil
}

func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case "ParseTopic":
		return engine.NewFuncParser(parser.ParseTopic, "ParseTopic"), nil
	case "ItemParser":
		if topic, ok := p.Args.(string); ok {
			return parser.NewItemParser(topic), nil
		} else {
			return nil, fmt.Errorf("invalid "+"arg: %v", p.Args)
		}
	case "NilParser":
		return engine.NilParser{}, nil
	default:
		return nil, errors.New("unknown parser name")
	}

}
func DeserializeResult(results []ParseResult) []engine.ParseResult {

	var array []engine.ParseResult
	for _, r := range results {
		result := engine.ParseResult{
			Items: r.Items,
		}
		fmt.Println(result)
		for _, req := range r.Requests {
			engineReq, err := DeserializeRequest(req)
			if err != nil {
				log.Printf("err deserialize request: %v", err)
				continue
			}
			result.Requests = append(result.Requests, engineReq)
		}
		array = append(array, result)
	}

	return array
}
