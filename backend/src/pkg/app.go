package pkg

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	"github.com/zencodeur/hoomi/backend/src/routes"
)

// App représente l'application Fiber avec ses middlewares
type App struct {
	*fiber.App
}

// NewApp crée une nouvelle instance de l'application avec les middlewares
func NewApp() *App {
	// Charger les variables d'environnement
	err := godotenv.Load()
	if err != nil {
		log.Println("Aucun fichier .env trouvé")
	}

	// Créer une nouvelle application Fiber
	app := fiber.New(fiber.Config{
		AppName: "Hoomi Backend API",
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin,Content-Type,Accept,Authorization",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))

	// Routes
	routes.SetupTestRoutes(app) // Passer l'app Fiber au lieu de router

	return &App{app}
}
