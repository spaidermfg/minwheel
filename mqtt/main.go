package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var ()

func main() {
	cat()
}

func cat() {
	// 创建 MQTT 客户端
	client := connect("client-1", "ssl://192.168.20.82:8883", onMessageReceived)

	// 订阅主题
	subscribe(client, "/DEVICE/BASE/MAGUS/cloud/#")

	// 等待消息接收
	waitForMessages()
}

func connect(clientID string, brokerURI string, onMessageReceived mqtt.MessageHandler) mqtt.Client {
	caCert, err := ioutil.ReadFile("/Users/mac/Downloads/ca.pem")
	if err != nil {
		panic(err)
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
	opts.SetDefaultPublishHandler(onMessageReceived)

	// 创建客户端实例
	client := mqtt.NewClient(opts)

	// 连接到 mqtt 服务器
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return client
}

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
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
