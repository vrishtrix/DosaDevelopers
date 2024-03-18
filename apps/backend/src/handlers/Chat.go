package handlers

import (
	"log"

	"dosadevelopers.devsoc/backend/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan models.Message)

func HomePage(c *fiber.Ctx) error {
	return c.SendString("Welcome to the Chat Room!")
}

func HandleConnections(c *websocket.Conn) {
	log.Println("WebSocket connection attempt")
	defer func() {
		if err := c.Close(); err != nil {
			log.Println("WebSocket could not be closed:", err)
		}
	}()

	clients[c] = true

	for {
		var msg models.Message
		if err := c.ReadJSON(&msg); err != nil {
			log.Println("Error reading json:", err)
			delete(clients, c)
			break
		}

		broadcast <- msg
	}
}

func HandleMessages() {
	for {
		msg := <-broadcast

		for client := range clients {
			if err := client.WriteJSON(msg); err != nil {
				log.Println("Error writing json:", err)
				if err := client.Close(); err != nil {
					log.Println("WebSocket could not be closed:", err)
				}
				delete(clients, client)
			}
		}
	}
}
