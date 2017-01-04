package main

import (
	// "fmt"
	"time"
	// MQTT "github.com/eclipse/paho.mqtt.golang"
	// "github.com/stianeikeland/go-rpio"
	"github.com/spf13/viper"
	"log"
)


func setupConfiguration() {
	viper.SetConfigName("garasjeagogo")
	viper.AddConfigPath("/etc/garasjeagogo/")
	viper.AddConfigPath("$HOME/.garasjeagogo/")
	viper.AddConfigPath("./")

	configError := viper.ReadInConfig()
	if configError != nil {
		log.Panicf("Error reading configuration: %v\n", configError)
	}

}


func main() {
	log.Println("Started garasjeagogo...")
	setupConfiguration()
	for {
		time.Sleep(30 * time.Second)
	}
}
