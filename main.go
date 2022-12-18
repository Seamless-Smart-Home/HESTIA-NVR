package main

import (
	"HESTIA/internal/compreface"
	"HESTIA/internal/database"
	"HESTIA/internal/frigate"
	frigateMqtt "HESTIA/internal/frigate/mqtt"
	mqttClient "HESTIA/internal/mqtt"
	"HESTIA/internal/scheduler"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	// Load Configurations using Viper
	LoadAppConfig()

	// Load the cron scheduler
	scheduler.Start()

	// Initialize Database
	err := database.Connect()
	if err != nil {
		panic(err)
	}
	database.Migrate()

	// Start MQTT Client
	mqttOptions := mqttClient.MqttOptions{
		Host:               AppConfig.MqttHost,
		Port:               AppConfig.MqttPort,
		Username:           AppConfig.MqttUsername,
		Password:           AppConfig.MqttPassword,
		MessagePubHandler:  messageHandler,
		ConnectHandler:     connectHandler,
		ConnectLostHandler: connectLostHandler,
	}
	err = mqttClient.BuildClient(mqttOptions)
	if err != nil {
		panic(err)
	}

	// Start Compreface Client
	err = compreface.NewClient(AppConfig.ComprefaceHost, AppConfig.ComprefaceAPIKey, nil)
	if err != nil {
		panic(err)
	}

	// Start Frigate Client
	err = frigate.Connect(AppConfig.FrigateHost)
	if err != nil {
		panic(err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Blocking, press ctrl+c to continue...")
	<-done // Will block here until user hits ctrl+c

	mqttClient.Client.Disconnect(250)

	// // Initialize Router
	// router := initRouter()
	// router.Run(fmt.Sprintf(":%v", AppConfig.ServerPort))

}

var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	topic := strings.Split(msg.Topic(), "/")

mqttTopic:
	switch topic[0] {
	case "frigate":
		switch topic[1] {
		case "events":
			err := frigateMqtt.ProcessEvent(msg.Payload())
			if err != nil {
				log.Println(err)
				break mqttTopic
			}
		default:
			if len(topic) >= 2 {
				switch topic[2] {
				case "person":
					switch topic[3] {
					case "snapshot":
						err := frigateMqtt.ProcessPersonDetected(msg.Payload())
						if err != nil {
							log.Println(err)
							break mqttTopic
						}
					}
				default:
					log.Println("Received unkown MQTT Topic: ", msg.Topic())
				}
			}
		}
	default:
		log.Println("Received unkown MQTT Topic: ", msg.Topic())
	}
	// fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

// func initRouter() *gin.Engine {
// 	router := gin.Default()
// 	api := router.Group("/api")
// 	{
// 		// Get Token Route
// 		api.POST("/token", controllers.GenerateToken)

// 		// Twilio Voice Routes
// 		voice := api.Group("/voice")
// 		{
// 			voice.POST("/answer", controllers.TwilioAnswer)
// 		}
// 	}
// 	return router
// }
