package main

import (
	"context"
	"fmt"
	pro "github.com/shirou/gopsutil/v4/process"
	"log"
	"os/exec"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("process start...")

	//name := "/Users/marksucik/Downloads/clipboard"
	//cmd := exec.CommandContext(context.Background(), name)
	//
	//if err := cmd.Start(); err != nil {
	//	log.Fatal("start error:", err)
	//}
	//
	//wg.Add(1)
	//c := make(chan struct{})
	//go stop(cmd, &wg, c)
	//
	//select {
	//case <-c:
	//	fmt.Println("process end...")
	//}
	//
	//wg.Wait()
	process()
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

func newCmdStop() {
	time.Sleep(6 * time.Second)
	name := "/Users/marksucik/Downloads/clipboard"
	cmd := exec.CommandContext(context.Background(), name)

	fmt.Println(cmd.Process.Pid)
	cmd.Process.Kill()
}

func process() {
	p, err := pro.NewProcess(33345)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p.Name())

	processes, err := pro.Processes()
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range processes {
		name, err1 := v.Name()
		if err1 == nil && name == "clipboard" {
			log.Println(v.IsRunning())
			log.Println(name, v.Pid)
		}
	}
}

func GetRunningStatus(name string) (bool, error) {
	p, err := pro.Processes()
	if err != nil {
		return false, err
	}

	for _, v := range p {
		n, err1 := v.Name()
		if err1 == nil && n == name {
			return v.IsRunning()
		}
	}

	return false, nil
}
