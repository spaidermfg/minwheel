package main

import (
	"flag"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
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

		n.search()
		n.processName()
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"序号", "协议", "本地地址", "外部地址", "连接状态", "进程ID", "进程名", "进程路径"})

		fmt.Println(">>> 输出信息开始渲染")
		for i, o := range n.Output {
			table.Append([]string{strconv.Itoa(i), o.Proto, o.LocalAddr, o.ForeignAddr, o.State, o.Pid, o.Name, o.Path})
		}
		fmt.Println(">>> 输出渲染完成")
		table.Render()
		os.Exit(0)
	}
	fmt.Println("请输入需要检索的进程相关信息\n\r[Example]:\n\r   [根据端口号检索: ./state.exe -tip 8200]\n\r   [根据进程名检索: ./state.exe -tip rtnn]")
}

func (n *netstat) search() {
	fmt.Println(">>> 开始查询端口使用信息")
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
		if len(fields) != 5 {
			continue
		}

		n.Output = append(n.Output, &output{
			Proto:       fields[0],
			LocalAddr:   fields[1],
			ForeignAddr: fields[2],
			State:       fields[3],
			Pid:         fields[4],
		})
	}
	fmt.Println(">>> 端口使用信息查询完成")
}

func (n *netstat) processName() {
	fmt.Println(">>> 通过PID开始查询进程信息")
	var wg sync.WaitGroup
	for _, v := range n.Output {
		wg.Add(1)
		go func(o *output) {
			defer wg.Done()
			cmd := exec.Command("powershell", "Get-Process", "-Id", o.Pid, "|", "Select-Object", "*")
			info, err := cmd.Output()
			if err != nil {
				log.Fatal(fmt.Sprintf("命令[%v]执行失败，Err: %v", cmd.String(), err))
			}

			res := strings.Split(string(info), "\n")
			for _, vi := range res {
				part := strings.SplitN(strings.TrimSpace(vi), ":", 2)
				if len(part) != 2 {
					continue
				}

				key := strings.TrimSpace(part[0])
				value := strings.TrimSpace(part[1])
				if key == "Path" {
					o.Path = value
				}
				if key == "ProcessName" {
					o.Name = value
				}
			}
		}(v)
	}

	wg.Wait()
	fmt.Println(">>> 进程信息查询完成")
}
