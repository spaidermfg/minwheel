package test

import (
	"fmt"
	jacoSerial "github.com/jacobsa/go-serial/serial"
	tarmSerial "github.com/tarm/serial"
	bugsSerial "go.bug.st/serial"
	"log"
	"testing"
)

func TestTarmSerial(t *testing.T) {
	c := &tarmSerial.Config{
		Name:        "/dev/tty.WI-XB400",
		Baud:        2400,
		ReadTimeout: 0,
		Size:        0,
		Parity:      0,
		StopBits:    0,
	}

	port, err := tarmSerial.OpenPort(c)
	if err != nil {
		log.Fatal("open port failed:", err)
	}
	defer port.Close()

	n, err := port.Write([]byte("hello world"))
	if err != nil {
		log.Fatal("write data failed:", err)
	}

	buf := make([]byte, 128)
	n, err = port.Read(buf)
	if err != nil {
		log.Fatal("read data failed:", err)
	}

	log.Println(buf[:n])
}

func TestBugsSerial(t *testing.T) {
	ports, err := bugsSerial.GetPortsList()
	if err != nil {
		log.Fatal("get serial port list failed: ", err)
	}

	if len(ports) == 0 {
		log.Fatal("No serial ports found.")
	}

	for k, v := range ports {
		fmt.Println(" num:", k, "| port:", v)
		fmt.Println("-------------------------------------------")
	}

	var flag bool

	if flag {
		mode := &bugsSerial.Mode{
			BaudRate:          0,
			DataBits:          0,
			Parity:            0,
			StopBits:          0,
			InitialStatusBits: nil,
		}

		port, err := bugsSerial.Open("", mode)
		if err != nil {
			log.Fatal("bugs open port failed:", err)
		}

		_, err = port.Write([]byte(""))
		if err != nil {
			log.Fatal("bugs write data failed:", err)
		}

		//port.
	}
}

func TestJacobsaSerial(t *testing.T) {
	option := jacoSerial.OpenOptions{
		PortName:                "",
		BaudRate:                0,
		DataBits:                0,
		StopBits:                0,
		ParityMode:              0,
		RTSCTSFlowControl:       false,
		InterCharacterTimeout:   0,
		MinimumReadSize:         0,
		Rs485Enable:             false,
		Rs485RtsHighDuringSend:  false,
		Rs485RtsHighAfterSend:   false,
		Rs485RxDuringTx:         false,
		Rs485DelayRtsBeforeSend: 0,
		Rs485DelayRtsAfterSend:  0,
	}

	//检查波特率是否正确
	//jacoSerial.IsStandardBaudRate()

	open, err := jacoSerial.Open(option)
	if err != nil {
		log.Fatal("Jaco open ports failed:", err)
	}
	defer open.Close()

	n, err := open.Write([]byte(""))
	if err != nil {
		log.Fatal("Jaco write data failed:", err)
	}

	buf := make([]byte, 128)
	n, err = open.Read(buf)
	if err != nil {
		log.Fatal("Jaco read data failed:", err)
	}

	log.Println(buf[:n])
}
