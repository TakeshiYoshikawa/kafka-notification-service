package main

import (
	"fmt"
	"kafka-notify/cmd/producers"
	"kafka-notify/pkg/models"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	users := []models.User{
		{ID: 1, Name: "Emma"},
		{ID: 2, Name: "Bruno"},
		{ID: 3, Name: "Richard"},
		{ID: 4, Name: "Leon"},
	}

	producer, err := producers.SetupProducer()

	if err != nil {
		log.Fatalf("Failed to initialize producer: %v", err)
	}
	defer producer.Close()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("/send", producers.SendMessageHandler(producer, users))

	fmt.Printf("Kafka producer -> Started at http://localhost%s\n", producers.ProducerPort)

	if err := router.Run(producers.ProducerPort); err != nil {
		log.Printf("Failed to run the server: %v", err)
	}
}
