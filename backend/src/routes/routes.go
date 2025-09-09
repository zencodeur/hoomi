package routes

import (
	"github.com/gofiber/fiber/v2" // Remplacer gin par fiber
)

// routes de test importées de test.go
func SetupTestRoutes(app *fiber.App) { // Modifier le type du paramètre
	app.Get("/test", TestRoute) // Utiliser app.Get au lieu de router.GET
}
