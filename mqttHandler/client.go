package mqttHandler

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/BwM17/mqtthook/config"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Init() {
	opts := mqttOptionsBuilder()
	client := mqtt.NewClient(opts)

	if config.Broker.Broker.Topic == "" {
		log.Fatalf("Error topic is required to run the program")
	}

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(fmt.Sprintf("Error connecting to MQTT broker: %v", token.Error()))
	}

	if token := client.Subscribe(config.Broker.Broker.Topic, 0, onMessageReceived); token.Wait() &&
		token.Error() != nil {
		panic(fmt.Sprintf("Error subscribing to topic: %v", token.Error()))
	}

	log.Println("Subsribed to topic:", config.Broker.Broker.Topic)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	client.Unsubscribe(config.Broker.Broker.Topic)
	client.Disconnect(250)
}
