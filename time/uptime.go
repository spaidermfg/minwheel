package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/host"
	"log"
	"time"
)

func main() {
	bootTime, err := host.BootTime()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("---------bootTime", bootTime)
	fmt.Println("----------since:", T(int64(bootTime)*1000))

	uptime, err := host.Uptime()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("---------uptime:", uptime)

	info, err := host.Info()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(info.String())
}

// T 毫秒时间戳计算
func T(t int64) string {
	startTime := time.UnixMilli(t)
	endTime := time.UnixMilli(time.Now().UnixMilli())
	duration := endTime.Sub(startTime)

	years := int(duration.Hours() / (24 * 365))
	duration -= time.Duration(years) * 365 * 24 * time.Hour

	months := int(duration.Hours() / (24 * 30))
	duration -= time.Duration(months) * 30 * 24 * time.Hour

	days := int(duration.Hours() / 24)
	duration -= time.Duration(days) * 24 * time.Hour

	hours := int(duration.Hours())
	duration -= time.Duration(hours) * time.Hour

	minutes := int(duration.Minutes())
	duration -= time.Duration(minutes) * time.Minute

	seconds := int(duration.Seconds())

	var output string
	if years > 0 {
		output += fmt.Sprintf("%d年", years)
	}
	if months > 0 {
		output += fmt.Sprintf("%d月", months)
	}
	if days > 0 {
		output += fmt.Sprintf("%d天", days)
	}
	if hours > 0 {
		output += fmt.Sprintf("%d小时", hours)
	}
	if minutes > 0 {
		output += fmt.Sprintf("%d分钟", minutes)
	}
	if seconds > 0 {
		output += fmt.Sprintf("%d秒", seconds)
	}

	return output
}
