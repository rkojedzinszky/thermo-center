package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/eclipse/paho.mqtt.golang"
)

type mqttClient struct {
	options *mqtt.ClientOptions
	hub     *Hub
	stop    chan struct{}
}

func newMqttClient(hub *Hub, mqttHost string, mqttPort int) *mqttClient {
	m := &mqttClient{
		options: mqtt.NewClientOptions(),
		hub:     hub,
		stop:    make(chan struct{}),
	}

	m.options.SetAutoReconnect(true)
	m.options.AddBroker(fmt.Sprintf("tcp://%s:%d", mqttHost, mqttPort))
	m.options.SetOnConnectHandler(func(cl mqtt.Client) {
		m.onConnect(cl)
	})

	return m
}

func (m *mqttClient) run() {
	client := mqtt.NewClient(m.options)

	client.Connect().Wait()

	<-m.stop

	client.Disconnect(0)

	close(m.hub.fromMqtt)
}

func (m *mqttClient) onConnect(cl mqtt.Client) {
	log.Print("Connected to MQTT broker")
	if cl.Subscribe("thsensor/+/report", 0, func(cl mqtt.Client, msg mqtt.Message) {
		m.onMessage(cl, msg)
	}).Wait() != true {
		log.Print("Subscribing failed")
	}
}

func (m *mqttClient) onMessage(cl mqtt.Client, msg mqtt.Message) {
	idstr := strings.Split(msg.Topic(), "/")
	var id uint64
	var err error
	if id, err = strconv.ParseUint(idstr[1], 16, 8); err != nil {
		log.Printf("Received invalid topic: %s\n", msg.Topic())
		return
	}

	m.hub.fromMqtt <- uint8(id)
}
