package test

import (
	"sync"
	"sync/atomic"
	"testing"
)

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
