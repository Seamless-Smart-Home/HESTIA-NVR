package mqttClient

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttOptions struct {
	Host               string `default:"localhost"`
	Port               int    `default:"1883"`
	Username           string
	Password           string
	ClientID           string `default:"hestia"`
	MessagePubHandler  mqtt.MessageHandler
	ConnectHandler     mqtt.OnConnectHandler
	ConnectLostHandler mqtt.ConnectionLostHandler
}

type MqttClient struct {
	mqttClient mqtt.Client
}

var Client *MqttClient

func BuildClient(options MqttOptions) error {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", options.Host, options.Port))
	opts.SetClientID(options.ClientID)
	opts.SetUsername(options.Username)
	opts.SetPassword(options.Password)
	opts.SetDefaultPublishHandler(options.MessagePubHandler)
	opts.OnConnect = options.ConnectHandler
	opts.OnConnectionLost = options.ConnectLostHandler
	mqttClient := mqtt.NewClient(opts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	c := &MqttClient{
		mqttClient: mqttClient,
	}

	Client = c

	return nil
}

func (c *MqttClient) Publish(topic string, message interface{}) {
	c.mqttClient.Publish(topic, 0, false, message)
}

func (c *MqttClient) Subscribe(topic string, handler mqtt.MessageHandler) {
	token := c.mqttClient.Subscribe(topic, 1, handler)
	fmt.Println("Subscribed to: ", topic)
	token.Wait()
}

func (c *MqttClient) Disconnect(delay uint) {
	c.mqttClient.Disconnect(delay)
}
