package config

import (
	"log"
	"os"
)

var Broker BrokerRoot
var Hook HookRoot

func LoadConfig() {
	yamlData, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("could not read config file: %v", err.Error())
	}

	loadYamlData(yamlData, &Broker, "broker", validateBroker)
	loadYamlData(yamlData, &Hook, "hook", validateHook)

}
