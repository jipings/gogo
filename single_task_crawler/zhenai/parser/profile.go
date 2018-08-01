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
)

func ParseProfile(contents []byte, name string) engine.ParseResult {
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
		Items: []interface{}{profile},
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
