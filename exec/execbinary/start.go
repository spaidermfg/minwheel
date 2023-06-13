package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	_, err := os.Stat("/opt/dasserver/server/magusdog")
	if os.IsNotExist(err) {
		log.Println("可执行文件不存在")
		return
	}

	err = exec.Command("/bin/sh", "-c", "./start.sh").Run()
	if err != nil {
		log.Println("magusdog执行失败")
		return
	}
}
