package test

import (
	"sync"
	"sync/atomic"
	"testing"
)

// ----------------------------------------
// 对共享整型变量的无锁读写
// ----------------------------------------

var n1 int64

func addByAtomic(i int64) int64 {
	return atomic.AddInt64(&n1, i)
}

func readByAtomic() int64 {
	return atomic.LoadInt64(&n1)
}

var n2 int64
var rwmu sync.RWMutex

func addByRWMutex(i int64) {
	rwmu.Lock()
	n2 += i
	rwmu.Unlock()
}

func readByRWMutex() int64 {
	var i int64
	rwmu.RLock()
	i = n2
	rwmu.RUnlock()
	return i
}

// go test -bench . atomic_test.go -cpu 2
// go test -bench . atomic_test.go -cpu 8
// go test -bench . atomic_test.go -cpu 16
// go test -bench . atomic_test.go -cpu 32
// 结论：
// 原子操作的无锁并发写随着并发量增大保持恒定
// 原子操作的无锁并发读的性能随着并发量增大性能逐渐提升，约为读锁的200倍
func BenchmarkAddByAtomic(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			addByAtomic(1)
		}
	})
}

func BenchmarkReadByAtomic(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			readByAtomic()
		}
	})
}

func BenchmarkAddByRWMutex(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			addByRWMutex(1)
		}
	})
}

func BenchmarkReadByRWMutex(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			readByRWMutex()
		}
	})
}

// ----------------------------------------
// 对共享自定义类型的读写
// 结论：
// 原子操作的无锁并发写的性能随着并发量增大而减小
// 原子操作的无所并发读的性能随着并发量增大而趋于稳定
// ----------------------------------------

type Config struct {
	sync.RWMutex
	data string
}

func BenchmarkRWMutexSet(b *testing.B) {
	config := Config{}
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			config.Lock()
			config.data = "easy"
			config.Unlock()
		}
	})
}

func BenchmarkRWMutexGet(b *testing.B) {
	config := Config{data: "language"}
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			config.RLock()
			_ = config.data
			config.RUnlock()
		}
	})
}

func BenchmarkAtomicSet(b *testing.B) {
	var config atomic.Value
	a := Config{data: "rust"}
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			config.Store(a)
		}
	})
}

func BenchmarkAtomicGet(b *testing.B) {
	var config atomic.Value
	config.Store(Config{data: "treaty"})
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = config.Load().(Config)
		}
	})
}
