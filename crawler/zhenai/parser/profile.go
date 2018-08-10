package parser

import (
	"regexp"
	"strconv"

	"imooc.com/tutorial/crawler/engine"
	"imooc.com/tutorial/crawler/model"
)

// 因为每次都重新编译正则会影响性能
// 因此直接在全局设置正则编译
// const (
// 	ageRe      = `<td><span class="label">年龄：</span>(\d+)岁</td>`
// 	marriageRe = `td><span class="label">婚况：</span>([^<]+)</td>`
// )

var (
	ageRe      = regexp.MustCompile(`<td><span class="label">年龄：</span>(\d+)岁</td>`)
	genderRe   = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^>]+)</span></td>`)
	heightRe   = regexp.MustCompile(`<td><span class="label">身高：</span><span field="">(\d+)CM</span></td>`)
	marriageRe = regexp.MustCompile(`td><span class="label">婚况：</span>([^<]+)</td>`)
	incomeRe   = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
	weightRe   = regexp.MustCompile(`<td><span class="label">身高：</span><span field="">(\d+)KG</span></td>`)
	guessRe    = regexp.MustCompile(`<a class="exp-user-name" [^>]*href="(http://album.zhenai.com/u/[\d]+)">([^<]+)</a>`)
	idurlRe    = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)
)

func ParseProfile(contents []byte, url, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}
	profile.Gender = extractString(contents, genderRe)
	profile.Marriage = extractString(contents, marriageRe)
	profile.Income = extractString(contents, incomeRe)

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai",
				Id:      extractString([]byte(url), idurlRe),
				Payload: profile,
			},
		},
	}

	// 猜你喜欢部分的解析
	mathches := guessRe.FindAllSubmatch(contents, -1)
	for _, m := range mathches {
		result.Requests = append(result.Requests,
			engine.Request{
				Url:    url,
				Parser: NewProfileParser(string(m[2])),
			},
		)
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}

	return ""

}

// 这个函数我有点疑惑了，参数还能分步骤传入啊，先传name，再
// 在函数返回时给新的函数传参，这是闭包的神奇之处啊！
// func ProfileParser(name string) engine.ParserFunc {
// 	return func(c []byte, url string) engine.ParseResult {
// 		return ParseProfile(c, url, name)
// 	}
// }
// 为了适应分布式爬虫需要重构ProfileParser
type ProfileParser struct {
	userName string
}

func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{
		userName: name,
	}
}

// 实现engine types中的Parser接口
func (p *ProfileParser) Parse(content []byte, url string) engine.ParseResult {
	return ParseProfile(content, url, p.userName)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return "ProfileParser", p.userName
}
