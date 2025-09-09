# Architecture du Backend Hoomi

## Structure des Dossiers

```
backend/
├── src/              # Code source principal
│   ├── main.go       # Point d'entrée de l'application
│   ├── controllers/  # Gestionnaires de requêtes
│   ├── middlewares/  # Middleware personnalisé
│   ├── models/       # Modèles de données
│   ├── pkg/          # Packages internes
│   │   └── app/      # Configuration de l'application (app.go)
│   ├── routes/       # Définition des routes
│   ├── services/     # Logique métier
│   └── utils/        # Fonctions utilitaires
├── go.mod            # Dépendances Go
├── go.sum            # Sommes de contrôle des dépendances
├── .env.example      # Exemple de variables d'environnement
└── README.md         # Documentation du backend
```

## Technologies

- Langage : Go
- Framework : Fiber v2
- Base de données : PostgreSQL
- Cache : Redis

## Packages

### pkg/app
Contient la configuration de l'application, y compris :
- Initialisation de Fiber
- Configuration des middlewares
- Gestion des variables d'environnement

### controllers
Gestionnaires de requêtes HTTP qui reçoivent les requêtes et renvoient les réponses.

### middlewares
Middleware personnalisé pour le traitement des requêtes (authentification, logging, etc.).

### models
Modèles de données représentant les entités de la base de données.

### routes
Définition et organisation des routes de l'API.

### services
Logique métier de l'application, indépendante du framework.

### utils
Fonctions utilitaires générales utilisées dans toute l'application.

## Documentation

- [Style API](adr-001-api-style.md)
- [Authentification](../security/authentication.md)
- [Gestion des Utilisateurs](../users/)

## Déploiement

- Docker
- CI/CD avec GitHub Actions