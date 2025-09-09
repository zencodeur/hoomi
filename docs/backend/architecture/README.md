# Architecture du Backend Hoomi

## Structure des Dossiers

```
backend/
├── api/              # Définitions des API
├── auth/             # Système d'authentification
├── models/           # Modèles de données
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
- [Gestion des Utilisateurs](users/)

## Déploiement

- Docker
- CI/CD avec GitHub Actions