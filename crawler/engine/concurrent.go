package engine

import (
	"log"
	"spiders_on_go/crawler/model"
)

// 定义了ConcurrentEngine结构体，包含了调度器接口和工作协程数量。
type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

// 定义了Scheduler接口，要求实现Submit和ConfigureMasterWorkerChan方法。
type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

// Run方法是并发引擎的核心，它首先创建请求和解析结果的通道，然后配置调度器，并为每个工作协程创建一个 worker 。
func (e *ConcurrentEngine) Run(seeds ...Request) {

	

	
  // 创建了两个通道，in用于传递请求，out用于接收解析结果。
	out := make(chan ParseResult)
	e.Scheduler.Run()
  
  // 根据指定的工作协程数量，创建相应数量的工作协程。
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(out, e.Scheduler)
	}


	// 将初始请求提交给调度器。
	for _, r := range seeds {
		if isDuplicate(r.Url) {
			//log.Printf("Duplicate request: "+
			//	"%s", r.Url)
			continue
		}
		e.Scheduler.Submit(r)
	}

	ProfileCount := 0

	for {
		result := <-out
		for _, item := range result.Items {
			if _, ok := item.(model.Profile); ok {
				log.Printf("Got profile #%d: %v", ProfileCount, item)
				ProfileCount++
			}
		}
		// URL 去重
		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				//log.Printf("Duplicate request: "+
				//	"%s", request.Url)
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

// 函数创建一个工作协程，不断从in通道接收请求，调用worker函数处理请求，并将结果发送到out通道
func createWorker(
	out chan ParseResult, s Scheduler) {
	in := make(chan Request)
	go func() {
		// 无限循环，从out通道接收解析结果，打印项目，并将新的请求提交给调度器。
		for {
			s.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}
