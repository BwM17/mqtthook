package config

//Broker
type BrokerRoot struct {
	Broker BrokerConfig `yaml:"broker"`
}

type BrokerConfig struct {
	Host     string `ymal:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Topic    string `yaml:"topic"`
}

//Hook
type HookRoot struct {
	Hook HookConfig `yaml:"hook"`
}

type HookConfig struct {
	Host    string `yaml:"host"`
	Method  string `yaml:"method"`
	Payload string `yaml:"payload"`
}
