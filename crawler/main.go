package main

import (
	"spiders_on_go/crawler/engine"
	"spiders_on_go/crawler/zhenai/parser"
)

func main() {

	engine.Run(engine.Request{
		Url:        "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
