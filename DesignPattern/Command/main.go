package main

import "fmt"

// 命令模式
// 将请求转化为一个包含与请求相关的所有信息的独立对象

// 发送者：触发命令
// 命令接口：声明一个执行命令的接口
// 具体命令：实现各种类型的请求
// 接受者：实现业务逻辑

// 适用场景：通过操作来参数化对象， 实现回滚恢复撤销功能

// Button 请求者
type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}

// Command 命令接口
type Command interface {
	execute()
}

// OnCommand 具体命令接口
type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() {
	c.device.on()
}

// OffCommand 具体命令接口
type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}

// Device 接收者接口
type Device interface {
	on()
	off()
}

type Tv struct {
	isRunning bool
}

func (t *Tv) on() {
	t.isRunning = true
	fmt.Println("TV is running.")
}

func (t *Tv) off() {
	t.isRunning = false
	fmt.Println("TV is shutdown.")
}

func main() {
	tv := &Tv{}

	onTv := &OnCommand{
		device: tv,
	}

	offTv := &OffCommand{
		device: tv,
	}

	onButton := &Button{
		command: onTv,
	}
	onButton.press()

	offButton := &Button{
		command: offTv,
	}
	offButton.press()
}
