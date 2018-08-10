package view

import (
	"os"
	"testing"

	"imooc.com/tutorial/crawler/engine"
	"imooc.com/tutorial/crawler/frontend/model"
	m "imooc.com/tutorial/crawler/model"
)

func TestSearchResult(t *testing.T) {
	out, err := os.Create("index.test.html")
	if err != nil {
		panic(err)

	}
	defer out.Close()
	view := CreateSearchResultView("index.html")

	page := model.SearchResult{
		Hits: 123,
	}

	item := engine.Item{
		Url:  "http://album.zhenai.com/u/108848769",
		Id:   "108848769",
		Type: "zhenai",
		Payload: m.Profile{
			Name:     "如果芸许",
			Gender:   "女",
			Age:      35,
			Height:   155,
			Income:   "3000元以下",
			Marriage: "离异"},
	}

	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}
	err = view.Render(out, page)
	if err != nil {
		panic(err)
	}

}
