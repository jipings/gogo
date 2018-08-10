package parser

import (
	"regexp"

	"imooc.com/tutorial/crawler/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, _ string) engine.ParseResult {

	reg := regexp.MustCompile(cityListRe)
	match := reg.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range match {
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
		})

	}
	return result

}
