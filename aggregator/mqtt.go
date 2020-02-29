package aggregator

import (
	"encoding/json"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const mqttTopicTemplate = "thsensor/%02x/report"

// MqttClient handles pushing to Mqtt
type MqttClient struct {
	options *mqtt.ClientOptions
	push    chan *mqttStat
}

type mqttStat struct {
	topic   string
	payload []byte
}

// NewMqttClient creates a new Mqtt Client
func NewMqttClient(mqttHost string, mqttPort int) *MqttClient {
	cl := &MqttClient{
		options: mqtt.NewClientOptions(),
		push:    make(chan *mqttStat, 1),
	}

	cl.options.AddBroker(fmt.Sprintf("tcp://%s:%d", mqttHost, mqttPort))
	cl.options.SetAutoReconnect(false)

	return cl
}

// Push enqueues a SensorStat to send
func (cl *MqttClient) Push(s sensorStat) bool {
	topic := fmt.Sprintf(mqttTopicTemplate, s.Sensor.Id)
	data, _ := json.Marshal(s.Stat)

	select {
	case cl.push <- &mqttStat{
		topic:   topic,
		payload: data,
	}:
		return true
	default:
		return false
	}
}

func (cl *MqttClient) connect() (mqtt.Client, error) {
	m := mqtt.NewClient(cl.options)
	token := m.Connect()
	token.Wait()
	if err := token.Error(); err != nil {
		return nil, err
	}

	return m, nil
}

func (cl *MqttClient) loop(m mqtt.Client) {
	for {
		ms := <-cl.push

		token := m.Publish(ms.topic, 0, false, ms.payload)
		token.Wait()
		if token.Error() != nil {
			break
		}
	}
}

// Run the mqtt loop
func (cl *MqttClient) Run() {
	for {
		m, _ := cl.connect()

		if m != nil {
			cl.loop(m)
			m.Disconnect(1000)
		}

		time.Sleep(1 * time.Second)
	}
}
