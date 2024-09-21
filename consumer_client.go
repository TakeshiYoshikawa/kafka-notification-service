package main

import (
	"context"
	"fmt"
	"kafka-notify/cmd/consumers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	store := &consumers.NotificationStore{
		Data: make(consumers.UserNotifications),
	}

	ctx, cancel := context.WithCancel(context.Background())
	go consumers.SetupConsumerGroup(ctx, store)
	defer cancel()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/notifications/:userID", func(ctx *gin.Context) {
		consumers.HandleNotifications(ctx, store)
	})

	fmt.Printf("Kafka Consumer (Group: %s) -> Started at http://localhost:%s\n", consumers.ConsumerGroup, consumers.ConsumerPort)

	if err := router.Run(consumers.ConsumerPort); err != nil {
		log.Printf("Failed to run the server: %v", err)
	}
}
