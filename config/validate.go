package config

import "errors"

func validateHook() (err error) {
	if Hook.Hook.Method != "GET" && Hook.Hook.Method != "POST" {
		return errors.New("Invalid request method for webhook")
	}

	if Hook.Hook.Host == "" {
		return errors.New("Hosturl for webhook required")
	}

	return nil
}

func validateBroker() (err error) {
	if Broker.Broker.Host != Broker.Broker.Host {
		return errors.New("Hosturl for MQTT broker required")
	}
	return nil
}
