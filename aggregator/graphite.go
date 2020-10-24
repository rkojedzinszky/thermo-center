package aggregator

import (
	"bytes"
	"fmt"
	"net"
	"text/template"
	"time"
)

// GraphiteSender handles pushing to Graphite
type GraphiteSender struct {
	addr               string
	push               chan []byte
	carbonPathTemplate *template.Template
}

// NewGraphiteSender creates a new Graphite sender
func NewGraphiteSender(host string, port int, pathTemplate string) *GraphiteSender {
	template, err := template.New("carbonPathTemplate").Parse(pathTemplate)
	if err != nil {
		return nil
	}

	return &GraphiteSender{
		addr:               fmt.Sprintf("%s:%d", host, port),
		push:               make(chan []byte, 1),
		carbonPathTemplate: template,
	}
}

type carbonMetric struct {
	SensorID int32
	Metric   string
}

// Push enqueues a SensorStat to send
func (g *GraphiteSender) Push(s sensorStat) bool {
	tstamp := fmt.Sprintf("%d", int(s.Sensor.LastTsf.Float64))

	buffer := bytes.NewBuffer(make([]byte, 0, 512))

	for m, v := range s.Stat {
		if fv, ok := v.(float64); ok {
			// Prepare carbon metric path
			g.carbonPathTemplate.Execute(buffer, &carbonMetric{
				SensorID: s.Sensor.ID,
				Metric:   m,
			})
			// append metric value, terminate line with newline
			fmt.Fprintf(buffer, " %f %s\n", fv, tstamp)
		}
	}

	select {
	case g.push <- buffer.Bytes():
		return true
	default:
		return false
	}
}

var dialer = net.Dialer{
	Timeout: 5 * time.Second,
}

func (g *GraphiteSender) connect() net.Conn {
	conn, err := dialer.Dial("tcp", g.addr)
	if err != nil {
		return nil
	}

	return conn
}

func (g *GraphiteSender) loop(c net.Conn) {
	for {
		data := <-g.push

		c.SetDeadline(time.Now().Add(5 * time.Second))
		n, err := c.Write(data)

		if err != nil || n != len(data) {
			break
		}
	}
}

// Run the mqtt loop
func (g *GraphiteSender) Run() {
	for {
		conn := g.connect()

		if conn != nil {
			g.loop(conn)
			conn.Close()
		}

		time.Sleep(1 * time.Second)
	}
}
