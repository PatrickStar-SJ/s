package engine

import (
	"log"
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
}

// Run方法是并发引擎的核心，它首先创建请求和解析结果的通道，然后配置调度器，并为每个工作协程创建一个 worker 。
func (e *ConcurrentEngine) Run(seeds ...Request) {

	// 创建了两个通道，in用于传递请求，out用于接收解析结果。
	in := make(chan Request)
	out := make(chan ParseResult)

	// 配置调度器，将in通道设置为调度器和工作协程之间的通信通道。
	e.Scheduler.ConfigureMasterWorkerChan(in)

	// 根据指定的工作协程数量，创建相应数量的工作协程。
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	// 将初始请求提交给调度器。
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item #%d: %v", itemCount, item)
			itemCount++
		}
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

// 函数创建一个工作协程，不断从in通道接收请求，调用worker函数处理请求，并将结果发送到out通道
func createWorker(
	in chan Request, out chan ParseResult) {
	go func() {
		// 无限循环，从out通道接收解析结果，打印项目，并将新的请求提交给调度器。
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
