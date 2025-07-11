package config

import (
	"log"

	"gopkg.in/yaml.v3"
)

// add validation function
func loadYamlData(in []byte, out interface{}, name string, validate func() error) {
	err := yaml.Unmarshal(in, out)
	if err != nil {
		log.Fatalf("Error loading %s config: %v", name, err.Error())
	}
	log.Printf("Successfully Loaded %s config\n", name)

	err = validate()
	if err != nil {
		log.Fatal(err.Error())
	}
}
