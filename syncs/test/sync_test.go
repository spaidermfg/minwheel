package test

import (
	"sync"
	"testing"
)

// 需要高性能的临界区同步机制场景
// channel属于高级同步原语，构建在低级同步原语上，性能比低级同步原语稍逊一筹
// 不要复制首次使用后的并发原语
// 并发量较小的情况下，互斥锁性能更好；读写锁适用于具有一定并发量且读多写少的场合。
var cs = 0
var mu sync.Mutex
var c = make(chan struct{}, 1)

func criticalSectionSyncByMutex() {
	mu.Lock()
	cs++
	mu.Unlock()
}

func criticalSectionSyncByChan() {
	c <- struct{}{}
	cs++
	<-c
}

// go test -bench . sync_test.go
func BenchmarkCriticalSectionSyncByMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		criticalSectionSyncByMutex()
	}
}

func BenchmarkCriticalSectionSyncByChan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		criticalSectionSyncByChan()
	}
}

// sync.Once实现单例模式
// 保证任意一个函数在程序运行期间只被执行一次

// sync.Pool 减轻垃圾回收压力 数据对象缓存池
// 建立临时缓存对象池
// var bufPool = sync.Pool{
// 	New: func() interface {
// 		return new(bytes.Buffer)
// 	},
// }
