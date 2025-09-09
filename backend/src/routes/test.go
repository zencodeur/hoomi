package routes

import (
	"github.com/gofiber/fiber/v2" // Remplacer gin par fiber
)

// routes de test
func TestRoute(c *fiber.Ctx) error { // Modifier le type du paramètre et retourner une erreur
	return c.JSON(fiber.Map{ // Utiliser c.JSON avec fiber.Map
		"message": "Route de test",
		"status":  "success",
	})
}
