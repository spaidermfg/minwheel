package test

import (
	"fmt"
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
