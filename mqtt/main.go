package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var ()

func main() {
	cat()
}

func cat() {
	// 创建 MQTT 客户端
	client, err := connect("client-1", "ssl://192.168.20.82:8883")
	if err != nil {
		panic(err)
	}

	// // 订阅主题
	// subscribe(client, "/DEVICE/BASE/MAGUS/edge/#")

	// // 等待消息接收
	// waitForMessages()

	c := subscribes(client, "/DEVICE/BASE/MAGUS/edge/#")
	go func() {
		for {
			select {
			case message := <-c:
				splitMes(message.Topic())
				fmt.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())
			}
		}
	}()
	waitForMessages()
}

func connect(clientID string, brokerURI string) (mqtt.Client, error) {
	caCert, err := ioutil.ReadFile("/Users/mac/Downloads/ca.pem")
	if err != nil {
		return nil, err
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
		RootCAs:            caCertPool,
	}

	opts := mqtt.NewClientOptions().
		AddBroker(brokerURI).
		SetClientID(clientID).
		SetCleanSession(true).
		SetTLSConfig(tlsConfig)

	// 注册消息接收回调函数
	//opts.SetDefaultPublishHandler(onMessageReceived)

	// 创建客户端实例
	client := mqtt.NewClient(opts)

	// 连接到 mqtt 服务器
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, err
	}

	return client, err
}

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	splitMes(message.Topic())
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())
}

func subscribe(client mqtt.Client, topic string) {
	if token := client.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Printf("Subscribed to topic: %s\n", topic)
}

func waitForMessages() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}

func splitMes(mes string) {
	split := strings.Split(strings.TrimPrefix(mes, "/"), "/")
	deviceCode := split[4]
	str := make([]string, 0)
	for i := 5; i < len(split); i++ {
		str = append(str, split[i])
	}
	topic := strings.Join(str, "/")
	fmt.Println("len: ", len(split), "code: ", deviceCode, "topic: ", topic)
}

func subscribes(client mqtt.Client, topic string) chan mqtt.Message {
	c := make(chan mqtt.Message)

	if token := client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		c <- msg
	}); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	return c
}
