package handlers

import (
	"log"
	"time"

	"dosadevelopers.devsoc/backend/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan models.Message)

// Predefined simple AI responses
var Replies = map[string]string{
	"hello":          "Hi wassup hwru?",
	"bye":            "bbye",
	"thanks":         "You're welcome!",
	"Where are you?": "I'm everywhere!",
	"How are you?":   "I'm good hby?",
}

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

		Response := generateResponse(msg.Message)
		time.Sleep(3 * time.Second)
		Msg := models.Message{
			Username: "Vispute",
			Message:  Response,
			Type:     "received",
		}

		broadcast <- Msg
	}
}

func HandleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			if err := client.WriteJSON(msg); err != nil {
				log.Println("Error writing json:", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func generateResponse(input string) string {
	for key, reply := range Replies {
		if input == key {
			return reply
		}
	}
	return "Hi"
}
