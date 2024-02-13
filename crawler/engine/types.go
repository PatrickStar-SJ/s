package engine

// 结构体包含了要抓取的 URL 和一个解析函数 ParserFunc
// ParserFunc函数的输入是一个字节切片（body,通常是网页的 HTML或者JSON格式）并返回一个 ParseResult 结构体实例。
// ParseResult实例包含了从body中解析出来的所有有用信息，这些信息被组织成Requests和Items两部分：
//   - Requests部分包含了从当前页面解析出来需要进一步访问的URL。这些URL将被爬虫系统加入到待抓取队列中，以便后续处理。
//   - Items部分包含了从当前页面解析出来的数据项。这些数据项可能是任何类型的数据，比如商品信息、文章内容等，它们是最终想要从网页中提取的信息。
//
// 通过这种方式，ParseResult允许爬虫系统在处理网页内容时保持高度的灵活性和扩展性
type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

// ParseResult则作为一个统一的接口，将解析的结果传递给爬虫系统的其他部分处理。
type ParseResult struct {
	Requests []Request     // 是一个 Request 切片，表示需要进一步抓取的 URL
	Items    []interface{} // 是一个空接口切片，可以存储任何类型的数据，通常用来存储解析出来的数据项。
}

func NilParser([]byte) ParseResult {
	return ParseResult{} // 不做任何处理，直接返回一个空的 ParseResult 结构体。
}
