package main

import (
	"fmt"
	"log"
	"os/exec"
)

type UnloadApplication struct{}

func main() {
	app := UnloadApplication{}
	app.GrepAllProcess()
}

// GrepAllProcess 查看所有相关正在运行的进程
func (u *UnloadApplication) GrepAllProcess() {
	cmd := exec.Command("ps")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", output)
}
