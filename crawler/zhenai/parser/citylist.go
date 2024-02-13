package parser

import (
	"regexp"
	"spiders_on_go/crawler/engine"
)

const cityListRe = `<a href="(.*www\.zhenai\.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

// 输入字节切片contents作为参数（这通常是网页的HTML内容），返回ParseResult类型
// 这个函数的作用是解析城市列表页面，提取出城市的URL和名称
// 这个函数是网络爬虫中负责解析城市列表的部分，是爬虫任务开始的一个入口点。
func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	// 在contents中查找所有匹配项。-1表示不限制查找的数量。
	//matches是一个二维字节切片，每个元素代表一个匹配项，其中第一个元素是整个匹配项，后续元素是各个捕获组的内容。
	matches := re.FindAllSubmatch(contents, -1)

	//初始化一个engine.ParseResult类型的变量result，用于存储解析结果。
	result := engine.ParseResult{}
	for _, m := range matches {
		//result.Items = append(
		//	result.Items, m[2])
		result.Items = append(
			result.Items, "City "+string(m[2])) //这里加了string之后，打印出城市名，否则会是 [int]的列表
		result.Requests = append(
			// 构造一个engine.Request对象，其中Url字段设置为城市的URL（转换为字符串），
			// ParserFunc字段设置为ParseCity（这是另一个函数，用于进一步解析每个城市的页面）。
			// 然后将这个Request对象追加到result.Requests切片中。
			result.Requests,
			engine.Request{
				Url:        string(m[1]),
				ParserFunc: ParseCity, //另一个函数，用于进一步解析每个城市的页面
			})
	}
	return result
}
