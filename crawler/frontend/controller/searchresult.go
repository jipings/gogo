package controller

import (
	"context"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	elastic "gopkg.in/olivere/elastic.v5"
	"imooc.com/tutorial/crawler/engine"
	"imooc.com/tutorial/crawler/frontend/model"
	"imooc.com/tutorial/crawler/frontend/view"
)

// TODO
// fill in query stringll in query string
// rewrite query string
// support search button
// support paging
// add start page
type SearchResultHandler struct {
	view   view.SearchResultView
	clinet *elastic.Client
}

// localhost:8888//searhc?q=男 & from=20
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))

	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}
	var page model.SearchResult
	page, err = h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	page.Query = q
	err = h.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)

	if err != nil {
		panic(err)
	}

	return SearchResultHandler{
		clinet: client,
		view:   view.CreateSearchResultView(template),
	}
}

func (h SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	resp, err := h.clinet.Search("profile").
		Query(elastic.NewQueryStringQuery(rewriterQueryString(q))).
		From(from).
		Do(context.Background())
	if err != nil {
		return result, err
	}

	result.Hits = int(resp.TotalHits())
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)
	return result, nil
}

func rewriterQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	res := re.ReplaceAllString(q, "Payload.$1:")
	log.Println(res)
	return res
}
