package scheduler

import "spiders_on_go/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s SimpleScheduler) WorkerReady(chan engine.Request) {

}

func (s *SimpleScheduler) Submit(
	r engine.Request) {
	//s.workerChan <- r   //会卡死，视频 16-3开头,16-2结尾。没看懂.换成下面开gorutine
	go func() { s.workerChan <- r }() // 关键是这一句：每个request会建一个gorutine，这个gorutine的作用就是往共用的channel去分发
}

func (s SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}
