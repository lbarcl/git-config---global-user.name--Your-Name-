package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func ReadConfig() Conf {
	data, err := os.ReadFile("./server.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var config Conf

	err = yaml.Unmarshal(data, &config)

	if err != nil {
		log.Fatal(err)
	}

	return config
}
