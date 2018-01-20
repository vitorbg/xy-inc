package conf

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	DatabaseName string `json:"database_name"`
	AppPort      int    `json:"app_port"`
}

var config Config

func LoadConf(path string) error {

	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetUrlDatabase() string {
	return config.DatabaseName
}

func GetAppPort() int {
	return config.AppPort
}
