package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type SensorData struct {
	DeviceID int     `json:"device_id"`
	Value    float64 `json:"value"`
}

func main() {

	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID("iot_device")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	defer client.Disconnect(250)

	var wg sync.WaitGroup
	numDevices := 500 


	for i := 0; i < numDevices; i++ {
		wg.Add(1)
		go func(deviceId int) {
			defer wg.Done()
			publishData(client, deviceId)
		}(i)
	}

	wg.Wait() 
}

func publishData(client mqtt.Client, deviceId int) {
	for i := 0; i < 10; i++ { 
		sensorData := SensorData{
			DeviceID: deviceId,
			Value:    rand.Float64() * 100,
		}

	
		sensorJsonData, err := json.Marshal(sensorData)
		if err != nil {
			log.Printf("Error marshaling JSON: %v", err)
			continue
		}
		token := client.Publish("iot/data", 0, false, sensorJsonData)
		token.Wait()
		log.Printf("Device %d published: %s", deviceId, string(sensorJsonData))
		time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
	}
}
