package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/namsral/flag"
)

func main() {
	mqttHost := flag.String("mqtt-host", "mqtt", "MQTT hostname/address")
	mqttPort := flag.Int("mqtt-port", 1883, "MQTT port")
	wsPort := flag.Int("ws-port", 8081, "Websocket port")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	hub := newHub()
	go hub.run(ctx)

	mqttclient := newMqttClient(hub, *mqttHost, *mqttPort)
	go mqttclient.run(ctx)

	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", *wsPort),
		Handler: hub,
	}

	go httpServer.ListenAndServe()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGTERM)
	<-sc
	cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Print(err)
	}
}
