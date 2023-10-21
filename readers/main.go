package main

import (
	"fmt"
	"github.com/chzyer/readline"
	"log"
	"strings"
)

var app = []string{"keyboard", "mouse", "video"}

func main() {
	reader, err1 := readline.New(">>> ")
	if err1 != nil {
		log.Fatal("create readline failed", err1)
	}
	defer reader.Close()

	reader.HistoryEnable()

	for {
		fmt.Print(">>> ")
		lines, err := reader.Readline()
		if err != nil {
			fmt.Println(err)
			continue
		}

		fields := strings.Fields(lines)
		f := len(fields)
		if f == 1 {
			if fields[0] == "exit" {
				return
			}

			if fields[0] == "ls" {
				for i, s := range app {
					fmt.Println(i, s)
				}
			}

			if fields[0] == "help" {
				fmt.Println("\n\t get some help information\n\t help\n\tExit the cli\n\t exit")
			}

			if fields[0] == "hello" {
				fmt.Println("world")
			}
		}

		if f > 1 {
			build := fields[0]
			branch := fields[1]
			plat := strings.Join(fields[2:], "")
			fmt.Printf("%v %v %v\n", build, branch, plat)
		}
	}
}
