package aggregator

import (
	"fmt"
	"net"
	"strings"
	"time"
)

// GraphiteSender handles pushing to Graphite
type GraphiteSender struct {
	addr string
	push chan []byte
}

// NewGraphiteSender creates a new Graphite sender
func NewGraphiteSender(host string, port int) *GraphiteSender {
	return &GraphiteSender{
		addr: fmt.Sprintf("%s:%d", host, port),
		push: make(chan []byte, 1),
	}
}

// Push enqueues a SensorStat to send
func (g *GraphiteSender) Push(s sensorStat) bool {
	tstamp := fmt.Sprintf("%d", int(s.Sensor.LastTsf.Float64))
	prefix := fmt.Sprintf("sensor.%02x", s.Sensor.Id)
	metrics := make([]string, 0, 10)

	for m, v := range s.Stat {
		if fv, ok := v.(float64); ok {
			metrics = append(metrics, fmt.Sprintf("%s.%s %f %s\n", prefix, m, fv, tstamp))
		}
	}

	data := []byte(strings.Join(metrics, ""))

	select {
	case g.push <- data:
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
