package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"
)

func main() {
	net_use()
	regular_use()
}

// 接收输入，校验IPv4地址
// net
func net_use() {
	var input string
	for {
		fmt.Println("请输入一个IP地址: ")
		fmt.Scanln(&input)
		ip := net.ParseIP(input)
		if ip == nil || ip.To4() == nil {
			fmt.Println("输入的IP的地址格式有误，请重新输入")
			continue
		}
		fmt.Println("输入的IP地址有效， IP为: ", ip)
		break
	}
}

// 正则表达式
func regular_use() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("请输入one个IP地址: ")
		ip, _ := reader.ReadString('\n')
		ip = strings.TrimSpace(ip)
		if isValidIP(ip) {
			fmt.Println("输入的IP地址有效-IP为: ", ip)
			break
		} else {
			fmt.Println("输入的IP的地址格式有误，请重新输入!")
		}
	}
}

func isValidIP(ip string) bool {
	// pattern := `^(\d{1,3}\.){3}\d{1,3}$`
	// matched, _ := regexp.MatchString(pattern, ip)

	ipRegex := regexp.MustCompile(`(\d{1,3}\.){3}\d{1,3}$`)
	flag := ipRegex.MatchString(ip)
	return flag
}
