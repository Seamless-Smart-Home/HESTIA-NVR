package main

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	FrigateHost      string `mapstructure:"FRIGATE_HOST"`
	MqttHost         string `mapstructure:"MQTT_HOST"`
	MqttPort         int    `mapstructure:"MQTT_PORT"`
	MqttUsername     string `mapstructure:"MQTT_USERNAME"`
	MqttPassword     string `mapstructure:"MQTT_PASSWORD"`
	ComprefaceHost   string `mapstructure:"COMPREFACE_HOST"`
	ComprefaceAPIKey string `mapstructure:"COMPREFACE_API_KEY"`
	ServerPort       string `mapstructure:"SERVER_PORT"`
}

var AppConfig *Config

func LoadAppConfig() {
	log.Println("Loading Server Configurations...")
	// Read file path
	viper.AddConfigPath(".")
	// set config file and path
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	// watching changes in app.env
	viper.AutomaticEnv()
	// reading the config file
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
}
