package broker

import (
	"fmt"
	"net/url"
	"os"
	"time"

	controller "ta2-script/controllers"
	logger "ta2-script/loggers"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// InitMQTT broker
func InitMQTT(cmdString string) {
	var mqttHost, mqttClient string
	if cmdString == "raw" {
		mqttHost = "MQTT_HOST_RAW"
		mqttClient = "MQTT_CLIENT_RAW"
	} else if cmdString == "finale" {
		mqttHost = "MQTT_HOST_FINALE"
		mqttClient = "MQTT_CLIENT_FINALE"
	}
	uri, err := url.Parse(os.Getenv(mqttHost))
	if err != nil {
		logger.LogRedError(err)
	}
	topic := uri.Path[1:len(uri.Path)]
	go listenMqtt(uri, mqttClient, topic)
	publishMqtt(uri, mqttClient, "trash_for_loop_"+topic)
}

// connectMqtt function
func connectMqtt(clientID string, uri *url.URL) mqtt.Client {
	opts := createClientOptions(clientID, uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		logger.LogRedError(err)
	}
	return client
}

// createClientOptions function
func createClientOptions(clientID string, uri *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	opts.SetPassword(password)
	opts.SetClientID(clientID)
	return opts
}

// listenMqtt function
func listenMqtt(uri *url.URL, clientID string, topic string) {
	client := connectMqtt("SUB-"+clientID, uri)
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
		if topic == "raw" {
			controller.CreateRawData(msg.Payload())
		} else if topic == "finale" {
			controller.CreateFinaleData(msg.Payload())
		}
	})
}

// publishMqtt function
func publishMqtt(uri *url.URL, clientID string, topic string) {
	client := connectMqtt("PUB-"+clientID, uri)
	timer := time.NewTicker(1 * time.Second)
	for t := range timer.C {
		client.Publish(topic, 0, false, t.String())
	}
}
