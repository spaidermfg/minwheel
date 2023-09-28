package test

import (
	"bytes"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

// 需要高性能的临界区同步机制场景
// channel属于高级同步原语，构建在低级同步原语上，性能比低级同步原语稍逊一筹
// 不要复制首次使用后的并发原语
// 并发量较小的情况下，互斥锁性能更好；读写锁适用于具有一定并发量且读多写少的场合。
// 互斥锁时临界区同步原语的首选，常用来对结构体对象的内部状态、缓存进行保护
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

// RWMutex 读写锁
// 读写锁的读锁性能恒定，不会随并发量的变化有巨大波动
// 并发量较大的情况下，读写锁的写锁性能较差，随着并发量的增大，性能有继续下降的趋势
// 适合应用在具有一定并发量且读多写少的场合
var cs1 = 0
var mu1 sync.Mutex
var cs2 = 0
var mu2 sync.RWMutex

// go test -bench .sync_test.go -cpu 2
func BenchmarkReadSyncByMutex(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu1.Lock()
			_ = cs1
			mu1.Unlock()
		}
	})
}

func BenchmarkReadSyncByRWMutex(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu2.RLock()
			_ = cs2
			mu2.RUnlock()
		}
	})
}

func BenchmarkWriteSyncByRWMutex(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu2.Lock()
			cs2++
			mu2.Unlock()
		}
	})
}

//Mutex等sync包中定义的结构类型在首次使用后不应对其进行复制操作
//由于首次使用后，Mutex进入locked状态，再次进行复制使用时，复制了locked状态的Mutex实例，导致程序进入死锁状态
//type Mutex struct {
//	state int32    //表示当前互斥锁的状态
//	sema uint32	   //用于控制锁状态的信号量
//}

// // sync.Once实现单例模式
// 保证任意一个函数在程序运行期间只被执行一次, 常用于初始化或资源清理过程中
type Foo struct {
	Name string
}

var once sync.Once
var instance *Foo

func GetInstance(id int) *Foo {
	defer func() {
		if e := recover(); e != nil {
			log.Printf("goroutine-%d: caught a panic: %s\n", id, instance)
		}
	}()

	log.Printf("goroutine-%d: enter GetInstance\n", id)
	once.Do(func() {
		instance = &Foo{Name: "mark"}
		time.Sleep(time.Second * 5)
		log.Printf("goroutine-%d: the addr of instance is %p\n", id, instance)
		panic("panic in once.Do function")
	})
	return instance
}

func TestOnceDo(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			getInstance := GetInstance(i)
			log.Printf("goroutine-%d: the addr of instance returned is %p\n", i, getInstance)
			wg.Done()
		}(i + 1)
	}
	time.Sleep(5 * time.Second)
	getInstance := GetInstance(0)
	log.Printf("goroutine-0: the addr of instance returned is %p \n", getInstance)
	wg.Wait()
	log.Println("all goroutine exit")
}

// sync.Pool 减轻垃圾回收压力 数据对象缓存池
// 并发安全的，放入缓存池中的数据对象是暂时的
// 建立临时缓存对象池
// 限制要放回缓存池中的数据对象大小, 不能超过64<<10
var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func writeBufFromPool(data string) {
	buffer := bufPool.Get().(*bytes.Buffer)
	buffer.Reset()
	buffer.WriteString(data)
	bufPool.Put(buffer)
}

func writeBufFromNew(data string) *bytes.Buffer {
	b := new(bytes.Buffer)
	b.WriteString(data)
	return b
}

// go test -bench . sync_test.go
func BenchmarkWithoutPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		writeBufFromPool("hello")
	}
}

func BenchmarkWithNew(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		writeBufFromNew("hello")
	}

	log.Println(64 << 10)
}

// sync.Map{}

var sm = sync.Map{}

func TestSyncMap(t *testing.T) {
	sm.Store("a", true)
	sm.Store("b", true)
	sm.Store("c", true)
	sm.Store("d", true)
	sm.Store("e", true)

	actual1, ok1 := sm.LoadOrStore("f", true)
	fmt.Println("ok1L", ok1)
	if ok1 {
		fmt.Printf("%v is exists\n", actual1)
	}

	if actual, ok := sm.LoadOrStore("a", true); ok {
		fmt.Printf("%v is exists\n", actual)
	}

	sm = sync.Map{}

	fmt.Println(sm.Load("b"))
	if actual, ok := sm.LoadOrStore("a", true); ok {
		fmt.Printf("%v is exists\n", actual)
	}
}
