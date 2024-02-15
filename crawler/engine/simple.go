package engine

import (
	"log"
	"spiders_on_go/crawler/fetcher"
)

type SimpleEngine struct{}

// 函数接收一个或多个 Request类型的参数 作为初始化种子，开始执行爬虫任务。
// 这里的“种子”并不是随机数生成中的种子概念，而是指启动爬虫过程的初始请求集.
// * 首先，它将所有传入的种子请求添加到一个本地的 requests 切片中
// * 然后，进入一个循环，只要 requests 切片不为空，就继续执行循环体。
func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	// 通过这个循环，爬虫引擎不断地抓取网页、解析内容，并根据解析结果中的新请求继续抓取，直到没有更多的请求需要处理。（requests切片，在循环里面也会扩充）
	// 这样，爬虫可以递归地抓取网站上的链接，收集和处理数据。
	for len(requests) > 0 {
		// 在循环体中，首先取出 requests 切片的第一个元素[0]，作为当前要处理的请求 r，并将其从切片中移除
		r := requests[0]
		requests = requests[1:]

		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		// 将解析结果中的新请求 parseResult.Requests 添加到 requests 切片中，以便后续继续抓取和解析
		requests = append(requests,
			parseResult.Requests...)
		// 遍历解析结果中的数据项 parseResult.Items，使用 log.Printf 打印每个数据项
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}

	}
}

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error "+
			"fetching url %s: %v",
			r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil

}

// Scheduler 实现1：所有worker共用一个输入
