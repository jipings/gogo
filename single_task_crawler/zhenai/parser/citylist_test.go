package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_contents.html")
	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents)
	const resultSize = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCity := []string{
		"City 阿坝", "City 阿克苏", "City 阿拉善盟",
	}
	if len(result.Requests) != resultSize {
		t.Errorf("[!] Request result should have %d requests; but had %d", resultSize, len(result.Requests))
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("[!] expected url #%d: %s; but was %s", i, url, result.Requests[i].Url)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("[1] Item result should have %d requests; but had %d", resultSize, len(result.Requests))
	}
	for i, city := range expectedCity {
		if result.Items[i].(string) != city {
			t.Errorf("[!] expected city $%d: %s; but was %s", i, city, result.Items[i].(string))
		}
	}

}
