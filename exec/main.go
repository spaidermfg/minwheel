package main

import (
	"fmt"
	"os/exec"
)

type UnloadApplication struct{}

func main() {
	app := UnloadApplication{}
	app.GrepAllProcess()
}

// GrepAllProcess 查看所有相关正在运行的进程
func (u *UnloadApplication) GrepAllProcess() {
	//cmd := exec.Command("bash", "-c", "ps -ef | awk '{print $2, $8}'")
	grepPid := "ps -def | grep dasserver | grep -v grep | awk '{print $1}'"
	output, err := exec.Command("/bin/sh", "-c", grepPid).Output()
	fmt.Println("=================@")
	if err != nil || len(output) != 0 {
		fmt.Println("=================&", len(output), err)
	}
	fmt.Println("=================#")
	fmt.Println(len(output))
}
