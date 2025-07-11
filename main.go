package main

import (
	"github.com/BwM17/mqtthook/config"
	"github.com/BwM17/mqtthook/mqttHandler"
)

func main() {
	config.LoadConfig()
	mqttHandler.Init()
}
