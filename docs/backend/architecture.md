# Architecture du Backend Hoomi

## Structure des Dossiers

```
backend/
├── controllers       # Contrôleurs pour les routes principales
│   ├── user.go       # Contrôleur pour les utilisateurs
│   └── ...           # Autres contrôleurs
├── middleware        # Middleware personnalisés
│   ├── auth.go       # Middleware d'authentification JWT
│   └── ...           # Autres middlewares
├── models/           # Modèles de données
├── routes            # Routes et contrôleurs
├── config/           # Configuration de l'application
├── services/         # Logique métier
├── utils/            # Fonctions utilitaires
└── ...
```

## Technologies

- Langage : Go
- Framework : Gin
- Base de données : PostgreSQL
- Cache : Redis

## Documentation

- [Style API](adr-001-api-style.md)
- [Authentification](../security/authentication.md)
- [Gestion des Utilisateurs](user-management.md)

## Déploiement

- Docker
- CI/CD avec GitHub Actions
