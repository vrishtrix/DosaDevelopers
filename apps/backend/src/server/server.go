package server

import (
	"os"
	"os/signal"
	"syscall"

	"dosadevelopers.devsoc/backend/src/config"
	"dosadevelopers.devsoc/backend/src/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/websocket/v2"
)

func StartServer() {
	handlers.ConnectToDB()
	defer handlers.DisconnectFromDB()
	// Fiber instance
	app := fiber.New()
	corsConfig := cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowMethods:     "POST, GET, OPTIONS, PUT, DELETE",
		AllowHeaders:     "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Ngrok-Skip-Browser-Warning",
		AllowCredentials: true,
		ExposeHeaders:    "Set-Cookie", // Include "Set-Cookie" in Access-Control-Expose-Headers
	}

	go handlers.HandleMessages()
	app.Use(cors.New(corsConfig))
	// Routes
	app.Use(logger.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	// Post Requests
	app.Post("/api/login", handlers.LoginHandler)
	app.Post("/api/signup", handlers.SignupHandler)

	// ChatApp
	app.Get("/ws", handlers.HomePage)
	app.Get("/ws/chat", websocket.New(handlers.HandleConnections))

	// Server Port
	err := app.Listen(":" + config.PORT)
	if err != nil {
		os.Exit(1)
	}

	// Check for termination signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Db disconnection Request
	handlers.DisconnectFromDB()
}
