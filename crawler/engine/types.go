package engine

// Parser统一接收content和url，
// 因为抽象后每个item都需要包含url,
type ParserFunc func(c []byte, url string) ParseResult

// added for distributed crawler
type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

// origin Request struct
// type Request struct {
// 	Url        string
// 	ParserFunc ParserFunc
// }
// modified for distributed crawler
type Request struct {
	Url    string
	Parser Parser
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Id      string
	Type    string // 在elastic search中type和id对应
	Payload interface{}
}

// origin NilParser，为了实现parser接口需要进行重构
// func NilParser([]byte) ParseResult {
// 	return ParseResult{}
// }
type NilParser struct{}

func (NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}
func (NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}

type FuncParser struct {
	parser ParserFunc
	name   string
}

// 所有的parser函数都将通过NewFuncParser创建
func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}

// NewFuncParser实现parser接口
func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}
func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}
