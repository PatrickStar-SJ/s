package parser

import (
	"regexp"
	"spiders_on_go/crawler/engine"
)

const cityRe = `<a href="(.*album\.zhenai\.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

// 这个函数用于解析城市页面
// 输入一个字节切片contents作为参数（这通常是城市页面的HTML内容）
// 这个函数的作用是解析城市页面，提取出用户的URL和名称。
func ParseCity(
	contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(
			result.Items, "User "+name)
		result.Requests = append(
			//构造一个engine.Request对象，其中Url字段设置为用户的URL（转换为字符串），
			//ParserFunc字段设置为一个匿名函数，该匿名函数调用ParseProfile函数并传入用户名。
			//然后将这个Request对象追加到result.Requests切片中。
			result.Requests,
			engine.Request{
				Url: string(m[1]),
				ParserFunc: func(c []byte) engine.ParseResult {
					return ParseProfile(c, name)
				},
			})
	}
	return result
}
