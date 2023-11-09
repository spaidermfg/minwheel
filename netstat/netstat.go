package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var tip = flag.String("tip", "", "输入进程相关信息")

type netstat struct {
	Tip    string
	Output []*output
}

type output struct {
	Proto       string
	LocalAddr   string
	ForeignAddr string
	State       string
	Pid         string

	Name string
	Path string
}

func main() {
	flag.Parse()

	if *tip != "" {
		n := netstat{
			Tip:    *tip,
			Output: nil,
		}

		fmt.Println("port", n.Tip)
		n.search()

		for i, o := range n.Output {
			fmt.Println(i, o.Proto, o.LocalAddr, o.ForeignAddr, o.State, o.Pid)
		}
		os.Exit(0)
	}
	fmt.Println("请输入需要检索的进程相关信息\n\r[Example]:\n\r   [根据端口号检索: ./pro 8200]\n\r   [根据进程名检索: ./pro rtnn]")
}

func (n *netstat) search() {
	cmd := exec.Command("cmd", "/C", fmt.Sprintf("netstat -ano | findstr %v", n.Tip))
	info, err := cmd.Output()
	if err != nil {
		log.Fatal(fmt.Sprintf("命令[%v]执行失败，Err: %v", cmd.String(), err))
	}

	lines := strings.Split(string(info), "\n")
	for _, line := range lines {
		fmt.Println(line)
		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) == 5 {
			n.Output = append(n.Output, &output{
				Proto:       fields[0],
				LocalAddr:   fields[1],
				ForeignAddr: fields[2],
				State:       fields[3],
				Pid:         fields[4],
			})
		}
	}
}

func (n *netstat) processName() {
	for _, o := range n.Output {
		cmd := exec.Command("cmd", "/C", fmt.Sprintf("Get-Process -Id %v | Select-Object *", o.Pid))
		//info, err := cmd.Output()
		if err != nil {
			log.Fatal(fmt.Sprintf("命令[%v]执行失败，Err: %v", cmd.String(), err))
		}

		//if err := json.Unmarshal(info); err != nil {
		//	return
		//}

	}
}
