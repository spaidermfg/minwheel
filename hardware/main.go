package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// disk
func main() {
	device := "sdb"
	lsblk := fmt.Sprintf("lsblk | grep disk | grep '%v'", device)
	cmd := exec.Command("/bin/sh", "-c", lsblk)
	output, err := cmd.Output()
	if err != nil {
		log.Println("err", err)
	}

	if strings.Contains(string(output), device) {
		fmt.Println("---------------------------------err")
	} else {
		fmt.Println("---------------------------------good")
	}

}
