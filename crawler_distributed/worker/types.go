package worker

import (
	"fmt"
	"log"

	"imooc.com/tutorial/crawler/engine"
	"imooc.com/tutorial/crawler/zhenai/parser"
	"imooc.com/tutorial/crawler_distributed/config"
)

// for distributed crawler
type SerializedParser struct {
	Name string
	Args interface{}
}

// 原来engine.Request中包含函数变量，
// 无法在网络中传递
type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

// 序列化engine.Request，转化为本地Request
func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}

}

// 序列化从engine.ParseResult，转化为本地ParseResult
func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result

}

func DeserializeRequest(r Request) (engine.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: parser,
	}, nil

}

func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		engineReq, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserializing request %v", err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}
	return result
}

func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList), nil
	case config.ParseCity:
		return engine.NewFuncParser(parser.ParseCity, config.ParseCity), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	case config.ParseProfile:
		if userName, ok := p.Args.(string); ok {
			return parser.NewProfileParser(userName), nil
		} else {
			return nil, fmt.Errorf("invalid arg: %v", p.Args)
		}
	default:
		return nil, fmt.Errorf("error deserailzing %v", p.Name)
	}
}
