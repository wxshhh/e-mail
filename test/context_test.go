package test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup
var exit bool

func TestContext(t *testing.T) {
	var err error
	fmt.Println(err)
	//ctx, cancel := context.WithCancel(context.Background())
	//fmt.Println(time.Duration(10000 * time.Second))
	//defer cancel()
	//for n := range gen(ctx) {
	//	fmt.Println(n)
	//	if n == 5 {
	//		break
	//	}
	//}
	// worker4
	//wg.Add(1)
	//ctx, cancel := context.WithCancel(context.Background())
	//go worker(ctx)
	//time.Sleep(3 * time.Second)
	//cancel()
	//wg.Wait()
	//fmt.Println("end")

	// worker3
	//wg.Add(1)
	//exitChan := make(chan struct{})
	//go Worker3(exitChan)
	//time.Sleep(3 * time.Second)
	//exitChan <- struct{}{}
	//close(exitChan)
	//wg.Wait()
	//fmt.Println("end")

	// worker2
	//wg.Add(1)
	//go Worker2()
	//time.Sleep(3 * time.Second)
	//exit = true
	//wg.Wait()
	//fmt.Println("end")

	// worker1
	//wg.Add(1)
	//go Worker1()
	//wg.Wait()
	//fmt.Println("end")
}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return // return结束该goroutine，防止泄露
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

func worker(ctx context.Context) {
	go worker2(ctx)
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}
	wg.Done()
}

func worker2(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker2")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}
}

func Worker4(ctx context.Context) {
	for {
		fmt.Println("worker4")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			wg.Done()
			return
		default:
		}
	}
}

func Worker3(exitChan chan struct{}) {
	for {
		fmt.Println("worker3")
		time.Sleep(time.Second)
		select {
		case <-exitChan:
			wg.Done()
			return
		default:
		}
	}
}

func Worker2() {
	for {
		fmt.Println("worker2")
		time.Sleep(time.Second)
		if exit {
			break
		}
	}
	wg.Done()
}

func Worker1() {
	for {
		fmt.Println("worker1")
		time.Sleep(time.Second)
	}
	wg.Done()
}
