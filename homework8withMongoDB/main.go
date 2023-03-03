package main

import (
    "anythings/configs" // Поменяйте на вашу папку
    "anythings/routes" // Поменяйте на вашу папку
    "github.com/gofiber/fiber/v2" 
)

func main() {
    app := fiber.New()

    configs.ConnectDB()
    
    routes.UserRoute(app)

    app.Listen(":8080")
}