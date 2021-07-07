package engine

type Request struct {
	Url       string // 请求地址
	ParseFunc func([]byte) (ParseResult, error) // 解析函数
}

type ParseResult struct {
	Requests []Request // 解析出的请求
	Items    []interface{} // 解析出的内容
}

// 实体数据
type Item struct {
	ID string
	URL string
	Type string
	PayLoad interface{} // 详细信息
}

func  NilParse([]byte) ParseResult {
	return  ParseResult{}
}