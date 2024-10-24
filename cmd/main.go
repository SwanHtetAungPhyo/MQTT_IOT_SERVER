package main

import (
	"log"

	"github.com/SwanHtetAungPhyo/IOT_Analysis/internal/services"
	"github.com/gofiber/fiber/v2"
)


func main(){
	go services.ConnectMQTT("tcp://localhost:1883", "iot/data")

	app := fiber.New()

	app.Get("/data", func(c *fiber.Ctx) error {
		data := services.GetData()
		return c.JSON(data)
	})

	services.PeriodicSave()
	log.Fatal(app.Listen(":8080"))
}