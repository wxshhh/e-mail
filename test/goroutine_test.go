package test

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	task := NewTask(func() error {
		fmt.Println("create a task:", time.Now().Format("2006-01-02 15:04:05"))
		return nil
	})
	pool := NewPool(3)
	go func() {
		for {
			pool.EntryChannel <- task
		}
	}()
	pool.Run()
}

type Task struct {
	f func() error
}

func NewTask(f func() error) *Task {
	return &Task{
		f: f,
	}
}

func (t *Task) Execute() {
	t.f()
}

type Pool struct {
	EntryChannel chan *Task
	JobChannel   chan *Task
	WorkerNum    int
}

func NewPool(workerNum int) *Pool {
	return &Pool{
		EntryChannel: make(chan *Task, workerNum),
		JobChannel:   make(chan *Task, workerNum),
		WorkerNum:    workerNum,
	}
}

func (p *Pool) worker(workID int) {
	for task := range p.JobChannel {
		task.Execute()
		fmt.Println("worker ", workID, " finish")
	}
}

func (p *Pool) Run() {
	for i := 0; i < p.WorkerNum; i++ {
		fmt.Println("start worker", i)
		go p.worker(i)
	}
	for task := range p.EntryChannel {
		p.JobChannel <- task
	}
	close(p.JobChannel)
	fmt.Println("worker finish")
	close(p.EntryChannel)
	fmt.Println("pool finish")
}
