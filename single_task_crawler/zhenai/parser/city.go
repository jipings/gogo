package parser

import (
	"regexp"

	"imooc.com/tutorial/crawler/engine"
)

const cityRe = `<th><a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^>]+)</a></th>`

func ParseCity(contents []byte) engine.ParseResult {

	reg := regexp.MustCompile(cityRe)
	match := reg.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range match {
		name := string(m[2]) // 防止闭包的延迟执行导致变量改变
		result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(contents []byte) engine.ParseResult {
				return ParseProfile(contents, name)
			},
		})
	}
	return result

}
