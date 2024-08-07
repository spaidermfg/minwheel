package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	pro "github.com/shirou/gopsutil/v4/process"
	"log"
	"os/exec"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("process start...")

	//name := "/Users/marksucik/Downloads/clipboard"
	//cmd := exec.CommandContext(context.Background(), name)
	//
	//if err := cmd.Start(); err != nil {
	//	log.Fatal("start error:", err)
	//}
	//
	//wg.Add(1)
	//c := make(chan struct{})
	//go stop(cmd, &wg, c)
	//
	//select {
	//case <-c:
	//	fmt.Println("process end...")
	//}
	//
	//wg.Wait()
	// process()
	diskInfo()
	cpuInfo()
}

func stop(cmd *exec.Cmd, wg *sync.WaitGroup, c chan struct{}) {
	defer wg.Done()
	fmt.Println("process cmd:", cmd.String())
	fmt.Println("process pid:", cmd.Process.Pid)
	ticker := time.NewTicker(6 * time.Second)
	select {
	case <-ticker.C:
		err := cmd.Process.Kill()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("process stop...")
		close(c)
	}
}

func process() {
	app := "clipboard"
	status, err := GetRunningStatus(app)
	if err != nil {
		log.Fatal("err1", err)
	}

	log.Println("status:", status)

	pid, err := GetProcessPid(app)
	if err != nil {
		log.Fatal("err2", err)
	}

	log.Println("pid:", pid)

	if err = StopProcess(context.Background(), app); err != nil {
		log.Fatal("err3: ", err)
	}

	status, err = GetRunningStatus(app)
	if err != nil {
		log.Fatal("err1", err)
	}

	log.Println("status:", status)

	pid, err = GetProcessPid(app)
	if err != nil {
		log.Fatal("err2", err)
	}

	log.Println("pid:", pid)
}

func GetRunningStatus(name string) (bool, error) {
	p, err := pro.Processes()
	if err != nil {
		return false, err
	}

	for _, v := range p {
		n, err1 := v.Name()
		if err1 == nil && n == name {
			return v.IsRunning()
		}
	}

	return false, nil
}

func StopProcess(ctx context.Context, name string) error {
	p, err := pro.Processes()
	if err != nil {
		return err
	}

	for _, v := range p {
		n, err1 := v.Name()
		if err1 == nil && n == name {
			log.Println(n, v.Pid)
			return v.KillWithContext(ctx)
		}
	}

	return errors.New("process not found")
}

func RestartProcess(ctx context.Context, name string) error {
	p, err := pro.Processes()
	if err != nil {
		return err
	}

	for _, v := range p {
		n, err1 := v.Name()
		if err1 == nil && n == name {
			log.Println(n, v.Pid)
			if err1 = v.KillWithContext(ctx); err1 != nil {
				return fmt.Errorf("stop process failed: %v", err1)
			}

			// start process
		}
	}

	return errors.New("process not found")
}

func GetProcessPid(name string) (int32, error) {
	p, err := pro.Processes()
	if err != nil {
		return 0, err
	}

	for _, v := range p {
		n, err1 := v.Name()
		if err1 == nil && n == name {
			return v.Pid, nil
		}
	}

	return 0, nil
}

func diskInfo() {
	d, err := disk.Usage("/")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(""+
		"total:%v\n"+
		"free:%v\n"+
		"percent:%v\n", d.Total/(1024*1024*1024), d.Free/(1024*1024*1024), d.UsedPercent)

	fmt.Println(d.String())
}

func cpuInfo() {
	//c, err := cpu.Info()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(c)
	//
	//percent, err := cpu.Percent(time.Second, false)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	////percentAll, err := cpu.Percent(time.Second, true)
	////if err != nil {
	////	log.Fatal(err)
	////}
	//
	//totalMhz := c[0].Mhz
	//usedMhz := totalMhz * percent[0] / 100.0
	//fmt.Println(totalMhz, usedMhz, percent, totalMhz/usedMhz)
	// 获取CPU信息
	info, err := cpu.Info()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(info)
	// 计算总的 MHz 和使用的 MHz
	var totalMHz float64
	var usedMHz float64
	for _, v := range info {
		totalMHz += v.Mhz
	}

	// 获取总体CPU使用率
	overallUsage, err := cpu.Percent(time.Second, false)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	usedMHz = totalMHz * (overallUsage[0] / 100.0)

	// 计算总的 CPU 使用率
	calculatedUsage := (usedMHz / totalMHz) * 100

	fmt.Printf("Total MHz: %.2f MHz, %v\n", totalMHz, totalMHz)
	fmt.Printf("Used MHz: %.2f MHz\n", usedMHz)
	fmt.Printf("Calculated CPU Usage: %.2f%%\n", calculatedUsage)
	fmt.Printf("Overall CPU Usage from cpu.Percent: %.2f%%\n", overallUsage[0])
}
