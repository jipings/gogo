package model

type SearchResult struct {
	Hits     int
	Query    string
	Start    int
	Items    []interface{}
	PrevFrom int
	NextFrom int
}
