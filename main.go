package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	server string = "127.0.0.1"
	topic  string = "hello"
	user   string = "user"
	passwd string = "root"
)

func onMessageReceived(client mqtt.Client, msg mqtt.Message) {
	log.Printf("Received message %s: from topic: %s\n", msg.Payload(), msg.Topic())

	payload, err := json.Marshal(map[string]string{
		"msg": "Im working",
	})
	if err != nil {
		log.Printf("Error marshalling JSON payload: %v", err)
		return
	}

	respBody := bytes.NewBuffer(payload)
	url := fmt.Sprintf("http://localhost:3000/%s", msg.Payload())
	fmt.Println(url)

	resp, err := http.Post(
		url,
		"application/json",
		respBody,
	)
	if err != nil {
		log.Printf("Error requesting %v", err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err.Error())
		return
	}

	sb := string(body)
	log.Println(sb)
}

func main() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("127.0.0.1:1883")
	opts.SetClientID("go_mqtt_client")
	opts.SetUsername(user)
	opts.SetPassword(passwd)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(fmt.Sprintf("Error connecting to MQTT broker:", token.Error()))
	}

	if token := client.Subscribe(topic, 0, onMessageReceived); token.Wait() &&
		token.Error() != nil {
		panic(fmt.Sprintf("Error subscribing to topic:", token.Error()))
	}

	log.Println("Subsribed to topic:", topic)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	client.Unsubscribe(topic)
	client.Disconnect(250)
}
