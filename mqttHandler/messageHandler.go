package mqttHandler

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/BwM17/mqtthook/config"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func postRequest(path string) (err error) {
	jsonStr := []byte(config.Hook.Hook.Payload)
	url := fmt.Sprintf("%v%v", config.Hook.Hook.Host, path)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return errors.New(fmt.Sprintf("Server responded with status code %v", res.StatusCode))
	}

	log.Printf("%v: success", res.Request.URL)

	_, err = io.Copy(io.Discard, res.Body)
	if err != nil {
		return err
	}

	return nil
}

func getRequest(path string) (err error) {
	url := fmt.Sprintf("%v%v", config.Hook.Hook.Host, path)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return errors.New(fmt.Sprintf("Server responded with status code %v", res.StatusCode))
	}

	log.Printf("%v: Success", res.Request.URL)

	_, err = io.Copy(io.Discard, res.Body)
	if err != nil {
		return err
	}

	return nil
}

func onMessageReceived(client mqtt.Client, msg mqtt.Message) {
	log.Printf("Recived message %s: from topic: %s\n", msg.Payload(), msg.Topic())

	switch config.Hook.Hook.Method {
	case "GET":
		err := getRequest(string(msg.Payload()))
		if err != nil {
			log.Printf("Error requestiong: %v", err.Error())
		}
		break
	case "POST":
		err := postRequest(string(msg.Payload()))
		if err != nil {
			log.Printf("Error requestiong: %v", err.Error())
		}
		break
	}
}
