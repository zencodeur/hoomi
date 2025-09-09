package main

import (
	"log"
	"github.com/zencodeur/hoomi/backend/src/pkg"
)

func main() {
	// Créer une nouvelle instance de l'application
	app := pkg.NewApp()


	// Démarrer le serveur
	log.Println("Démarrage du serveur Hoomi Backend...")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Impossible de démarrer le serveur: %v", err)
	}
}