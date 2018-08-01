package parser

import (
	"regexp"

	"imooc.com/tutorial/crawler/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {

	reg := regexp.MustCompile(cityListRe)
	match := reg.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	limit := 10
	for _, m := range match {
		result.Items = append(result.Items, "City "+string(m[2]))
		limit--
		if limit == 0 {
			break
		}
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})

	}
	return result

}
