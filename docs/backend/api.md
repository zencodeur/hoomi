# Spécification de l'API Backend Hoomi

Ce document décrit l'API REST du backend de Hoomi.

## Informations Générales

- **Version de l'API** : v1
- **Format de réponse** : JSON
- **Protocole** : HTTPS
- **Authentification** : Bearer Token (JWT)

## Endpoints

### Authentification

#### `POST /api/v1/auth/register`

- **Description** : Enregistre un nouvel utilisateur.
- **Corps de la requête** :
  ```json
  {
    "email": "string",
    "password": "string",
    "name": "string"
  }
  ```
- **Réponses** :
  - `201 Created` : Utilisateur créé avec succès.
    ```json
    {
      "user": {
        "id": "string",
        "email": "string",
        "name": "string"
      },
      "token": "string"
    }
    ```
  - `400 Bad Request` : Données invalides.
  - `409 Conflict` : L'email est déjà utilisé.

#### `POST /api/v1/auth/login`

- **Description** : Authentifie un utilisateur.
- **Corps de la requête** :
  ```json
  {
    "email": "string",
    "password": "string"
  }
  ```
- **Réponses** :
  - `200 OK` : Authentification réussie.
    ```json
    {
      "user": {
        "id": "string",
        "email": "string",
        "name": "string"
      },
      "token": "string"
    }
    ```
  - `400 Bad Request` : Données invalides.
  - `401 Unauthorized` : Identifiants incorrects.

### Utilisateurs

#### `GET /api/v1/users/me`

- **Description** : Récupère les informations de l'utilisateur actuel.
- **En-têtes** : `Authorization: Bearer <token>`
- **Réponses** :
  - `200 OK` :
    ```json
    {
      "id": "string",
      "email": "string",
      "name": "string"
    }
    ```
  - `401 Unauthorized` : Token invalide ou manquant.

#### `PUT /api/v1/users/me`

- **Description** : Met à jour le profil de l'utilisateur.
- **En-têtes** : `Authorization: Bearer <token>`
- **Corps de la requête** :
  ```json
  {
    "name": "string"
  }
  ```
- **Réponses** :
  - `200 OK` :
    ```json
    {
      "id": "string",
      "email": "string",
      "name": "string"
    }
    ```
  - `400 Bad Request` : Données invalides.
  - `401 Unauthorized` : Token invalide ou manquant.

### Hoomis (Groupes Familiaux)

#### `POST /api/v1/hoomis`

- **Description** : Crée un nouveau Hoomi.
- **En-têtes** : `Authorization: Bearer <token>`
- **Corps de la requête** :
  ```json
  {
    "name": "string"
  }
  ```
- **Réponses** :
  - `201 Created` :
    ```json
    {
      "id": "string",
      "name": "string",
      "createdAt": "string"
    }
    ```
  - `400 Bad Request` : Données invalides.
  - `401 Unauthorized` : Token invalide ou manquant.

#### `GET /api/v1/hoomis`

- **Description** : Liste les Hoomis de l'utilisateur.
- **En-têtes** : `Authorization: Bearer <token>`
- **Réponses** :
  - `200 OK` :
    ```json
    [
      {
        "id": "string",
        "name": "string",
        "createdAt": "string"
      }
    ]
    ```
  - `401 Unauthorized` : Token invalide ou manquant.

#### `GET /api/v1/hoomis/{id}`

- **Description** : Récupère les détails d'un Hoomi.
- **En-têtes** : `Authorization: Bearer <token>`
- **Paramètres** : `id` (string) - ID du Hoomi.
- **Réponses** :
  - `200 OK` :
    ```json
    {
      "id": "string",
      "name": "string",
      "members": [
        {
          "id": "string",
          "name": "string"
        }
      ],
      "createdAt": "string"
    }
    ```
  - `401 Unauthorized` : Token invalide ou manquant.
  - `403 Forbidden` : L'utilisateur n'est pas membre du Hoomi.
  - `404 Not Found` : Hoomi non trouvé.

#### `POST /api/v1/hoomis/{id}/invite`

- **Description** : Invite un utilisateur dans un Hoomi.
- **En-têtes** : `Authorization: Bearer <token>`
- **Paramètres** : `id` (string) - ID du Hoomi.
- **Corps de la requête** :
  ```json
  {
    "email": "string"
  }
  ```
- **Réponses** :
  - `200 OK` : Invitation envoyée.
  - `400 Bad Request` : Données invalides.
  - `401 Unauthorized` : Token invalide ou manquant.
  - `403 Forbidden` : L'utilisateur n'a pas le droit d'inviter.
  - `404 Not Found` : Hoomi non trouvé.

### Messages

#### `POST /api/v1/hoomis/{id}/messages`

- **Description** : Envoie un message dans un Hoomi.
- **En-têtes** : `Authorization: Bearer <token>`
- **Paramètres** : `id` (string) - ID du Hoomi.
- **Corps de la requête** :
  ```json
  {
    "content": "string"
  }
  ```
- **Réponses** :
  - `201 Created` :
    ```json
    {
      "id": "string",
      "content": "string",
      "senderId": "string",
      "createdAt": "string"
    }
    ```
  - `400 Bad Request` : Données invalides.
  - `401 Unauthorized` : Token invalide ou manquant.
  - `403 Forbidden` : L'utilisateur n'est pas membre du Hoomi.
  - `404 Not Found` : Hoomi non trouvé.

#### `GET /api/v1/hoomis/{id}/messages`

- **Description** : Récupère les messages d'un Hoomi.
- **En-têtes** : `Authorization: Bearer <token>`
- **Paramètres** : `id` (string) - ID du Hoomi.
- **Paramètres de requête** :
  - `limit` (integer, optionnel) : Nombre de messages à récupérer (par défaut 50).
  - `before` (string, optionnel) : ID du message avant lequel récupérer les messages (pour la pagination).
- **Réponses** :
  - `200 OK` :
    ```json
    [
      {
        "id": "string",
        "content": "string",
        "senderId": "string",
        "senderName": "string",
        "createdAt": "string"
      }
    ]
    ```
  - `401 Unauthorized` : Token invalide ou manquant.
  - `403 Forbidden` : L'utilisateur n'est pas membre du Hoomi.
  - `404 Not Found` : Hoomi non trouvé.

## Gestion des Erreurs

Toutes les erreurs de l'API renvoient un objet JSON avec la structure suivante :
```json
{
  "error": {
    "code": "string",
    "message": "string"
  }
}
```