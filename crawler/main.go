package main

import (
	"spiders_on_go/crawler/engine"
	"spiders_on_go/crawler/scheduler"
	"spiders_on_go/crawler/zhenai/parser"
)

/*
我想知道是先执行的parser.ParseCityList，还是先执行的engine.Run？

	parser.ParseCityList并不是被执行的函数，而是作为一个参数（具体来说，是一个函数引用或函数指针）传递给了engine.Run函数。
	engine.Run函数是先被执行的，它根据自己的逻辑来决定何时执行传递给它的parser.ParseCityList函数。
	这种设计允许engine模块控制整个爬虫的流程，而具体的页面解析逻辑则可以灵活地通过不同的解析函数实现，增强了代码的模块化和可扩展性。
*/
func main() {

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		Url:        "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
