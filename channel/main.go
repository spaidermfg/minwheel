package main

import (
	"fmt"
	"log"
	"minwheel/channel/chant"
	"minwheel/channel/webbug"
	"os"
	"sync"

	"gopl.io/ch8/thumbnail"
)

//channel是goroutine之间互相通信的通道
//通过通信来共享内存
//无缓冲的通道(同步通道)：chan := make(chan int)，发送和接收必须同步
//创建带缓存区的通道： channel := make(chan int, 10)
//创建的是数据结构的引用，零值是nil
//发送： send := make(chan<- int)
//接收: receive := <-chan
//关闭： close(ch)
//chan<- int: 这个类型只能发送
//<-chan int: 这个类型只能接收
//关闭一个已经关闭的通道会宕机，关闭一个仅能接受的通道编译时报错，关闭通道后发送操作会导致宕机
//goroutine泄漏：发送消息到通道但没有goroutine来接收
//-- 带缓冲的channel
//发送操作相缓存队列尾部插入元素，接收操作从队列的头部删除元素
//使用cap获取channel内部缓存的容量
//使用len获取channel内部缓存队列中有效元素的个数

func main() {
	// c := make(chan int, 1)
	// go send(c)
	// go receive(c)
	// time.Sleep(3 * time.Second)
	// close(c)

	// pipline()
	// time.Sleep(10 * time.Second)

	// naturals := make(chan int)
	// squares := make(chan int)
	// go counter(naturals)
	// go squarer(squares, naturals)
	// printer(squares)

	chant.Chant()
	handleImage()
  webbug.Bug()
}

func send(c chan<- int) {
	for i := 0; i <= 10; i++ {
		c <- i
		fmt.Println("send: ", i)
	}
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println("received: ", v)
	}
}

func tcpClient() {

}

func pipline() {
	naturals := make(chan int)
	squares := make(chan int)
	go func() {
		for i := 0; i <= 100; i++ {
			naturals <- i
		}
		close(naturals)
	}()

	go func() {
		for v := range naturals {
			squares <- v * v
		}
		close(squares)
	}()

	go func() {
		for v := range squares {
			fmt.Println("+++++++", v)
		}
	}()
}

func counter(out chan<- int) {
	for i := 0; i <= 100; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

type item struct {
	thumbfile string //filename
	err       error  //error
	size      int64  //file size
}

// 并发的循环
// 循环处理图片
func makeThumbnails(filenames []string) (thumbfiles []item, err error) {
	//创建通道计数
	//ch := make(chan struct{})
	ch := make(chan item, len(filenames))
	var wg sync.WaitGroup //工作的goroutine个数
	for _, v := range filenames {
		wg.Add(1)
		// if _, err := thumbnail.ImageFile(v); err != nil {
		// 	log.Println(err)
		// }
		go func(v string) {
			defer wg.Done()
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(v)
			fi, _ := os.Stat(it.thumbfile)
			it.size = fi.Size()
			log.Println(it.size)
			ch <- it

		}(v)
	}

	go func() {
		//等待和关闭
		wg.Wait()
		close(ch)
	}()

	//wait goroutine exec finish
	for range filenames {
		it := <-ch
		if it.err != nil {
			log.Println(err)
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it)
	}
	return thumbfiles, nil
}

func handleImage() {
	a := []string{"a.jpg"}
	thumbfiles, err := makeThumbnails(a)
	if err != nil {
		log.Fatal("[error]: ", err)
	}

	for _, v := range thumbfiles {
		log.Println("| FileName: ", v.thumbfile, "| FileSize: ", v.size, "| error: ", v.err)
	}
}
