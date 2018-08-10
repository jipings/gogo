package parser

import (
	"regexp"

	"imooc.com/tutorial/crawler/engine"
)

var (
	profileRe = regexp.MustCompile(`<th><a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^>]+)</a></th>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)">下一页</a>`)
)

func ParseCity(contents []byte, _ string) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	// 此处代码与profile.go中的有重复，需要进行重构
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: NewProfileParser(string(m[2])),
		})
	}
	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests,
			engine.Request{
				Url:    string(m[1]),
				Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
			},
		)
	}
	return result

}
