package main

import (
	"fmt"
	"time"
	"net/http"
	_ "net/http/pprof"
)

type Info struct {
	Params interface{}
	Func   func(interface{})
}

type Task struct {
	info chan Info
}

type Pool struct {
	tasks *Task
	num   int
}

func NewPool(n int) *Pool {
	taskChan := make(chan Info)
	return &Pool{
		tasks: &Task{
			info: taskChan,
		},
		num: n,
	}
}

func (p *Pool) Put(info Info) {
	p.tasks.info <- info
}

func (p *Pool) Run() {
	for gc := 0; gc < p.num; gc++ {
		go p.Work(gc)
	}
}

func (p *Pool) Work(workId int) {
	for {
		select {
		case job := <-p.tasks.info:
			params := fmt.Sprintf("param:%s workid:%d", job.Params, workId)
			job.Func(params) //执行传入的函数
		}
	}
}

func Proc(params interface{}) {
	fmt.Println("Result:", params)
}

func main() {

	p := NewPool(4)
	go p.Run()
	go http.ListenAndServe("0.0.0.0:6060", nil) // pprof查看协程使用情况
	for {
		p.Put(Info{Params: "http://www.boyaa.com", Func: Proc})
		time.Sleep(time.Millisecond * 500)
	}
}
