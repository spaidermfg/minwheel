package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type Network struct {
	Businfo string `json:"businfo"`
	Physid  string `json:"physid"`
	Serial  string `json:"serial"`
}

type rule struct {
	Name string
	Attr string
}

type Rules struct {
	Data []*rule
}

var show = flag.Bool("show", false, "查看网口信息")
var write = flag.Bool("write", false, "生成udev网口规则文件")
var reload = flag.Bool("reload", false, "保存udev网口配置规则\nsudo udevadm control --reload-rules")

func main() {
	flag.Parse()

	if *show {
		r := new(Rules)
		network := r.analysisStdout()
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"网口名称", "mac地址"})
		r.getNetInfo(network)
		for _, v := range r.Data {
			table.Append([]string{v.Name, v.Attr})
		}
		table.Render()
		os.Exit(0)
	}

	if *write {
		r := new(Rules)
		network := r.analysisStdout()
		r.getNetInfo(network)
		r.writeRulesFile()
		os.Exit(0)
	}

	if *reload {
		cmd := exec.CommandContext(context.Background(), "udevadm", "control", "--reload-rules")
		if _, err1 := cmd.Output(); err1 != nil {
			log.Fatal("sudo udevadm control --reload-rules执行失败:", err1)
		}
		fmt.Println("保存成功")
		os.Exit(0)
	}

	fmt.Println("请输入-h选项获取帮助")
}

// 获取排序好的网口信息
func (r *Rules) getNetInfo(data []*Network) {
	add := func(name, attr string) {
		r.Data = append(r.Data, &rule{
			Name: name,
			Attr: attr,
		})
	}

	for _, v := range data {
		if v.Businfo == "" {
			switch v.Physid {
			case "1":
				add("lan1", v.Serial)
			case "2":
				if !r.exist5G() {
					add("lan2", v.Serial)
				}
			case "3":
				if r.exist5G() {
					add("lan2", v.Serial)
				}
			}
		} else {
			switch v.Businfo {
			case "pci@0000:04:00.0":
				add("lan3", v.Serial)
			case "pci@0000:04:00.1":
				add("lan4", v.Serial)
			case "pci@0000:04:00.2":
				add("lan5", v.Serial)
			case "pci@0000:04:00.3":
				add("lan6", v.Serial)
			case "pci@0000:02:00.0":
				add("lan7", v.Serial)
			case "pci@0000:02:00.1":
				add("lan8", v.Serial)
			case "pci@0000:02:00.2":
				add("lan9", v.Serial)
			case "pci@0000:02:00.3":
				add("lan10", v.Serial)
			}
		}
	}

	sort.Slice(r.Data, func(i, j int) bool {
		a := strings.TrimPrefix(r.Data[i].Name, "lan")
		b := strings.TrimPrefix(r.Data[j].Name, "lan")

		ai, err := strconv.Atoi(a)
		if err != nil {
			log.Fatal("数据类型有误:", err)
		}

		bi, err := strconv.Atoi(b)
		if err != nil {
			log.Fatal("数据类型有误:", err)
		}
		return ai < bi
	})
}

// 生成udev网络规则临时文件
func (r *Rules) writeRulesFile() {
	file, err := os.OpenFile("10-rename-network-temp.rules", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Fatal("创建文件有误:", err)
	}

	l := strings.Builder{}
	for _, v := range r.Data {
		l.WriteString(fmt.Sprintf("SUBSYSTEM==\"net\", ACTION==\"add\", DRIVERS==\"?*\", ATTR{address}==\"%v\", NAME=\"%v\"\n", v.Attr, v.Name))
	}

	if _, err = file.WriteString(l.String()); err != nil {
		log.Fatal("文件写入失败:", err)
	}

	abs, err := filepath.Abs(file.Name())
	if err != nil {
		log.Fatal("文件路径获取失败:", err)
	}

	fmt.Println("生成网口规则临时文件成功 | Path:", abs)
}

// 查看是否存在有方的5G芯片
func (r *Rules) exist5G() bool {
	cmd := exec.CommandContext(context.Background(), "bash", "-c", "lsusb | grep Neoway")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal("lsusb 命令执行失败:", err)
	}

	if string(output) != "" {
		return true
	}

	return false
}

// 解析lshw命令输出
func (r *Rules) analysisStdout() []*Network {
	cmd := exec.Command("lshw", "-json", "-C", "network")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal("lshw -json -C network命令执行失败:", err)
	}

	var network []*Network
	if err = json.Unmarshal(output, &network); err != nil {
		log.Fatal("文件内容有误:", err)
	}

	return network
}
