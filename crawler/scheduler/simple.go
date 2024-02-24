package scheduler

import "spiders_on_go/crawler/engine"

// 定义了SimpleScheduler结构体，包含一个工作协程通道。
type SimpleScheduler struct {
	workerChan chan engine.Request
}

// Submit方法用于提交请求，它通过一个goroutine将请求发送到工作协程通道，避免直接发送导致的阻塞。
func (s *SimpleScheduler) Submit(
	r engine.Request) {
	//s.workerChan <- r   //会卡死，视频 16-3开头,16-2结尾。没看懂.换成下面开gorutine
	go func() { s.workerChan <- r }() // 关键是这一句：每个request会建一个gorutine，这个gorutine的作用就是往共用的channel去分发
}

// 用于配置主工作协程通道，实际上就是将调度器内部的工作协程通道设置为外部传入的通道。
func (s *SimpleScheduler) ConfigureMasterWorkerChan(
	c chan engine.Request) {
	s.workerChan = c
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}
