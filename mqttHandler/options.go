package mqttHandler

import (
	"fmt"
	"log"

	"github.com/BwM17/mqtthook/config"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func validateOpts() {}

func mqttOptionsBuilder() (options *mqtt.ClientOptions) {
	opts := mqtt.NewClientOptions()

	if config.Broker.Broker.Host != "" && config.Broker.Broker.Port != 0 {
		log.Printf("%v:%v\n", config.Broker.Broker.Host, config.Broker.Broker.Port)
		opts.AddBroker(fmt.Sprintf("%v:%v", config.Broker.Broker.Host, config.Broker.Broker.Port))
	}

	if config.Broker.Broker.Host != "" && config.Broker.Broker.Port == 0 {
		log.Printf("%v:%v\n", config.Broker.Broker.Host, 1883)
		opts.AddBroker(fmt.Sprintf("%v:%v", config.Broker.Broker.Host, 1883))
	}

	if config.Broker.Broker.Username != "" {
		opts.SetUsername(config.Broker.Broker.Username)
	}

	if config.Broker.Broker.Password != "" {
		opts.SetPassword(config.Broker.Broker.Password)
	}

	return opts
}
