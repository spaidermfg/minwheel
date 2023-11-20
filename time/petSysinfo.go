package main

import (
	"log"
	"os/exec"
	"sync"
	"time"
)

type command struct {
	Cmd  string
	Arg1 string
	Arg2 string
}

const cmdName = "pet_sysinfo"

func main() {
	onGPIO4, offGPIO4 := newCommand("-l 0", "on", "off")
	onGPO1, offGPO1 := newCommand("-x", "0101", "0000")
	onGPO2, offGPO2 := newCommand("-x", "1010", "0000")

	light := make(map[*command]*command)
	light[onGPIO4] = offGPIO4
	light[onGPO1] = offGPO1
	light[onGPO2] = offGPO2

	var wg sync.WaitGroup
	for on, off := range light {
		wg.Add(1)
		go func(on *command, off *command) {
			defer wg.Done()
			on.execCmd()
			off.execCmd()
		}(on, off)
	}

	wg.Wait()
}

func newCommand(args, onArgs, offArgs string) (*command, *command) {
	on := &command{
		Cmd:  cmdName,
		Arg1: args,
		Arg2: onArgs,
	}

	off := &command{
		Cmd:  cmdName,
		Arg1: args,
		Arg2: offArgs,
	}
	return on, off
}

func (c *command) execCmd() {
	cmd := exec.Command(c.Cmd, c.Arg1, c.Arg2)
	if err := cmd.Run(); err != nil {
		log.Println("命令", cmd.String(), "执行失败", err)
	}
}

func (c *command) toggleLight() {
	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.execCmd()
		}
	}
}
