package services

import (
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	mqttClient mqtt.Client
	memoryData []any
	mu         sync.RWMutex
)


func ConnectMQTT(broker string, topic string) {
	opts := mqtt.NewClientOptions().AddBroker(broker)
	opts.SetClientID("mqtt_server")
	mqttClient := mqtt.NewClient(opts)

	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	if token := mqttClient.Subscribe(topic, 0, messageHandler); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
}


func messageHandler(client mqtt.Client, msg mqtt.Message) {
	log.Printf("Received message: %s from topic: %s", msg.Payload(), msg.Topic())


	payload := string(msg.Payload())

	mu.Lock()
	memoryData = append(memoryData, payload)
	mu.Unlock()
}

// Saves the memoryData to a JSON file
func saveToJson(memoryData []any) {
	jsonData, err := json.MarshalIndent(memoryData, "", "	")
	if err != nil {
		log.Fatal(err.Error())
	}

	file, err := os.Create("./output.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	_, err = file.Write(jsonData)

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Data saved to output.json")
}


func GetData() []any {
	mu.Lock()
	defer mu.Unlock()
	return memoryData
}


func PeriodicSave() {
	go func() {
		for {
			time.Sleep(30 * time.Second) 
			mu.Lock()
			saveToJson(memoryData)
			mu.Unlock()
			log.Println("Data saved periodically to output.json")
		}
	}()
}
