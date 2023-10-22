package main

import (
	"context"
	"fmt"
	"github.com/chzyer/readline"
	"log"
	"os/exec"
	"strings"
	"sync"
	"time"
)

type Deploy struct {
	host *Host
	task *Task

	date time.Time
	ch   chan struct{}
	ctx  context.Context
	wg   *sync.WaitGroup
	mu   sync.Mutex
}

type Task struct {
	name     string
	property string
	absPath  string
	port     string
	branch   string
}

type Host struct {
	host     string
	port     string
	user     string
	arch     string
	password string
}

const (
	// cmd
	ls      = "ls" // ls task, ls
	cd      = "cd"
	gen     = "gen"
	run     = "run"
	help    = "help"
	down    = "down" // download task to local
	exit    = "exit"
	build   = "build"
	history = "history"
	push    = "push" // option: store,remote

	// file property
	text   = "text"
	binary = "binary"

	// arch
	arm   = "arm"
	amd64 = "amd64"
	arm64 = "arm64"
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
					stdout()
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

func stdout() func() {
	cmd := exec.Command("ls")
	bytes, err := cmd.Output()
	fmt.Println(string(bytes), err)

	return func() {
		fmt.Println("return func")
	}
}
