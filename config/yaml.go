package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var instance *Conf

func ReadConfig() Conf {
	if instance == nil {
		fmt.Println("Reading config file")
		data, err := os.ReadFile("./server.yaml")
		if err != nil {
			log.Fatal(err)
		}

		var config Conf

		err = yaml.Unmarshal(data, &config)

		if err != nil {
			log.Fatal(err)
		}

		instance = &config
		return config
	}
	fmt.Println("Returning config file")
	return *instance

}
