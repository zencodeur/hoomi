# Système de Gestion des Utilisateurs Hoomi

## Vue d'Ensemble

Ce document décrit le système de gestion des utilisateurs dans Hoomi, qui permet aux membres de créer, modifier et gérer leurs profils et comptes.

## Fonctionnalités

### Création de Compte

- Formulaire d'inscription avec email, nom et mot de passe
- Vérification de l'unicité de l'email
- Validation des critères de mot de passe
- Envoi d'email de confirmation (à implémenter)

### Profil Utilisateur

#### Informations de Base
- Nom complet
- Photo de profil
- Date de naissance
- Bio courte

#### Paramètres
- Préférences de notification
- Langue de l'interface
- Thème (clair/sombre)
- Confidentialité du profil

### Gestion du Compte

#### Sécurité
- Changement de mot de passe
- Gestion des sessions actives
- Authentification à deux facteurs (2FA)
- Historique des connexions

#### Données Personnelles
- Export des données utilisateur
- Suppression de compte
- RGPD : droit à l'oubli

### Administration

#### Pour les Admins de Hoomi
- Visualisation des membres
- Modification des rôles
- Suppression de membres

## Architecture

### Modèles de Données

#### Utilisateur (User)
- `id`: UUID
- `email`: String (unique)
- `password_hash`: String
- `name`: String
- `profile_picture_url`: String (optionnel)
- `date_of_birth`: Date (optionnel)
- `bio`: String (optionnel)
- `language`: String (par défaut 'fr')
- `theme`: Enum ['light', 'dark'] (par défaut 'light')
- `is_email_verified`: Boolean (par défaut false)
- `created_at`: Timestamp
- `updated_at`: Timestamp

#### Profil Utilisateur (UserProfile)
- `user_id`: UUID (clé étrangère vers User)
- `notification_preferences`: JSON
- `privacy_settings`: JSON
- `last_login_at`: Timestamp
- `last_activity_at`: Timestamp

### API Endpoints

#### Authentification
- `POST /api/auth/register`
- `POST /api/auth/login`
- `POST /api/auth/logout`

#### Gestion du Profil
- `GET /api/users/me`
- `PUT /api/users/me`
- `PUT /api/users/me/password`
- `DELETE /api/users/me`

#### Administration (Hoomi)
- `GET /api/hoomis/:hoomiId/members`
- `PUT /api/hoomis/:hoomiId/members/:userId/role`
- `DELETE /api/hoomis/:hoomiId/members/:userId`

## Sécurité

### Protection des Données
- Hachage des mots de passe avec bcrypt
- Validation des entrées utilisateur
- Protection contre les attaques CSRF

### RGPD
- Export complet des données utilisateur sur demande
- Suppression totale des données personnelles
- Consentement explicite pour les traitements de données

## Intégration avec d'Autres Systèmes

### Authentification
- Intégration avec le système d'authentification JWT
- Synchronisation avec le système de sessions

### Médias
- Stockage des photos de profil via le système de fichiers sécurisé
- Génération de miniatures

### Notifications
- Envoi de notifications pour les événements liés au compte
- Paramétrage des préférences de notification

## Évolution Future

### Fonctionnalités Planifiées
- Vérification d'email par lien de confirmation
- Récupération de mot de passe
- Connexion via réseaux sociaux
- Profils personnalisés par Hoomi
- Badges et réalisations