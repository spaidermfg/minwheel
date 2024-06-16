package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("process start...")

	name := "/Users/marksucik/Downloads/clipboard"
	cmd := exec.CommandContext(context.Background(), name)

	if err := cmd.Start(); err != nil {
		log.Fatal("start error:", err)
	}

	wg.Add(1)
	c := make(chan struct{})
	go stop(cmd, &wg, c)

	select {
	case <-c:
		fmt.Println("process end...")
	}

	wg.Wait()
}

func stop(cmd *exec.Cmd, wg *sync.WaitGroup, c chan struct{}) {
	defer wg.Done()
	fmt.Println("process cmd:", cmd.String())
	fmt.Println("process pid:", cmd.Process.Pid)
	ticker := time.NewTicker(6 * time.Second)
	select {
	case <-ticker.C:
		err := cmd.Process.Kill()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("process stop...")
		close(c)
	}
}
