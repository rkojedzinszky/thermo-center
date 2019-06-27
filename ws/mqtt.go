package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

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

	m.options.SetConnectTimeout(5 * time.Second)
	m.options.SetMaxReconnectInterval(10 * time.Second)
	m.options.AddBroker(fmt.Sprintf("tcp://%s:%d", mqttHost, mqttPort))
	m.options.SetOnConnectHandler(func(cl mqtt.Client) {
		m.onConnect(cl)
	})

	return m
}

func (m *mqttClient) run() {
	var client mqtt.Client

	// Try connecting
	// Once connected, mqtt will keep reconnecting
	for {
		client = mqtt.NewClient(m.options)
		token := client.Connect()
		token.Wait()
		if token.Error() == nil {
			break
		}

		// Handle stop request during just connection attempts
		select {
		case <-m.stop:
			close(m.hub.fromMqtt)
			return
		default:
		}
	}

	<-m.stop

	close(m.hub.fromMqtt)

	client.Disconnect(0)
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
