package test

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCh(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 8
	a := <-ch
	fmt.Println(a)

	s := make(chan struct{})
	go func(s chan struct{}) {
		time.Sleep(5 * time.Second)
		close(s)
	}(s)

	<-s
}

func TestTs(t *testing.T) {
	tick := time.NewTicker(time.Second * 2)
	ctx, cancel := context.WithCancel(context.Background())

	defer func() {
		tick.Stop()
		cancel()
	}()

	m := func(wg *sync.WaitGroup, ch chan<- any) {
		defer wg.Done()
		status, err := getPointStatus()
		if err != nil {
			ch <- err
			return
		}
		ch <- status
	}

	p := func(wg *sync.WaitGroup, ch chan<- any) {
		defer wg.Done()
		info, err := getPointInfo(ctx, "b")
		if err != nil {
			ch <- err
			return
		}
		ch <- info
	}

	s := func(wg *sync.WaitGroup, ch chan<- any) {
		defer wg.Done()
		monitorInfo, err := getMonitorInfo(ctx, "a")
		if err != nil {
			ch <- err
			return
		}
		ch <- monitorInfo
	}

	for {
		select {
		case <-tick.C:
			now := time.Now()
			ch := make(chan any, 4)

			wg := new(sync.WaitGroup)
			wg.Add(3)
			go m(wg, ch)
			go p(wg, ch)
			go s(wg, ch)
			wg.Wait()
			close(ch)

			res := new(Result)
			for v := range ch {
				switch r := v.(type) {
				case error:
					fmt.Println(r)
					res.E = errors.Join(res.E, r)
				case *Monitor:
					fmt.Println(r)
					res.M = r
				case *Point:
					fmt.Println(r)
					res.P = r
				case bool:
					fmt.Println(r)
					res.CacheRunningState = r
				}
			}

			fmt.Println(time.Since(now), res)
		}
	}
}

type Result struct {
	CacheRunningState bool // 缓存库运行状态
	M                 *Monitor
	P                 *Point
	E                 error
}

type Monitor struct {
	CpuPercent      float64
	MemoryPercent   float32
	RunningState    bool // 服务运行状态
	DataSourceState bool // 数据源运行状态

	CacheRunningState bool  // 缓存库运行状态
	LastTime          int64 // 最后启动时间
}

func getMonitorInfo(ctx context.Context, name string) (*Monitor, error) {
	time.Sleep(time.Second * 4)
	if name == "a" {
		return &Monitor{
			CpuPercent:        10,
			MemoryPercent:     20,
			RunningState:      true,
			DataSourceState:   false,
			CacheRunningState: false,
			LastTime:          50,
		}, nil
	} else {
		return nil, errors.New("monitor not found")
	}
}

type Point struct {
	Good    int
	Bad     int
	Timeout int
	Total   int
}

func getPointInfo(ctx context.Context, name string) (*Point, error) {
	time.Sleep(time.Second * 2)
	if name == "a" {
		return &Point{
			Good:    20,
			Bad:     1,
			Timeout: 30,
			Total:   230,
		}, nil
	} else {
		return nil, errors.New("point info err")
	}
}

func getPointStatus() (bool, error) {
	time.Sleep(time.Second * 1)
	return true, nil
}

func TestCh1(t *testing.T) {
	ch := make(chan int, 2)

	go sendCh(ch)
	select {
	case a := <-ch:
		fmt.Println("----", a)
	}
}

func sendCh(ch chan int) {
	time.Sleep(time.Second * 2)
	ch <- 1

	time.Sleep(time.Second * 2)
	ch <- 2
}
