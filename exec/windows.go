package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

const (
	BinaryName = "dass-backup.exe"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("getwd: ", err)
	}
	fmt.Println("Path:", dir)

	cmd := exec.Command(BinaryName, "list")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal("cmd.output():", err)
	}

	fmt.Println(string(output))
}
