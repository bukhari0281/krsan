package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"krsan/configs"
	"krsan/routes"
	"log"
	"os"
)

func BootApp() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var portEnv string

	if portEnv = os.Getenv("PORT"); portEnv == "8000" {
		portEnv = "8080"
	}

	configs.BootDatabase()
	configs.ConnectDatabase()
	configs.RunMigration()

	app := fiber.New()
	// CORS Config
	app.Use(cors.New(cors.Config{
		AllowOrigins:     configs.AllowOrigins,
		AllowMethods:     configs.AllowMethods,
		AllowHeaders:     configs.AllowHeaders,
		AllowCredentials: configs.AllowCredentials,
		ExposeHeaders:    configs.ExposeHeaders,
		MaxAge:           configs.MaxAge,
	}))

	// init route
	routes.InitRoute(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"hello": "Dunia",
		})
	})

	app.Listen(":8000")
}
