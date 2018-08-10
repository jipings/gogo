package engine

import (
	"log"

	"imooc.com/tutorial/crawler/fetcher"
)

func Worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(r.Url)

	if err != nil {
		log.Printf("Fetch: error fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	log.Printf("Fetch: fetching url %s", r.Url)
	parseResult := r.Parser.Parse(body, r.Url)
	return parseResult, nil

}
