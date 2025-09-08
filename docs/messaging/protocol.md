# Protocole de Messagerie Hoomi

Je documente ici le protocole de messagerie que j'ai défini pour Hoomi, qui gère à la fois les échanges via l'API REST et les communications en temps réel via WebSockets.

## Messages API REST

### Envoi d'un Message

**Requête :**
```
POST /api/v1/hoomis/{hoomiId}/messages
Authorization: Bearer <JWT>
Content-Type: application/json

{
  "content": "<contenu_chiffré_en_base64>",
  "nonce": "<nonce_utilisé_pour_le_chiffrement>",
  "sessionId": "<identifiant_de_la_session_de_chiffrement>"
}
```

**Réponse Succès :**
```
201 Created
Content-Type: application/json

{
  "id": "uuid-du-message",
  "hoomiId": "uuid-du-hoomi",
  "senderId": "uuid-de-l-expediteur",
  "content": "<contenu_chiffré_en_base64>",
  "nonce": "<nonce_utilisé_pour_le_chiffrement>",
  "sessionId": "<identifiant_de_la_session_de_chiffrement>",
  "createdAt": "2023-10-27T10:00:00Z"
}
```

### Récupération des Messages

**Requête :**
```
GET /api/v1/hoomis/{hoomiId}/messages?limit=50&before=uuid-message-cursor
Authorization: Bearer <JWT>
```

**Réponse Succès :**
```
200 OK
Content-Type: application/json

[
  {
    "id": "uuid-du-message-1",
    "hoomiId": "uuid-du-hoomi",
    "senderId": "uuid-de-l-expediteur-1",
    "senderName": "Nom de l'expéditeur 1",
    "content": "<contenu_chiffré_en_base64>",
    "nonce": "<nonce_utilisé_pour_le_chiffrement>",
    "sessionId": "<identifiant_de_la_session_de_chiffrement>",
    "createdAt": "2023-10-27T10:00:00Z"
  },
  {
    "id": "uuid-du-message-2",
    "hoomiId": "uuid-du-hoomi",
    "senderId": "uuid-de-l-expediteur-2",
    "senderName": "Nom de l'expéditeur 2",
    "content": "<contenu_chiffré_en_base64>",
    "nonce": "<nonce_utilisé_pour_le_chiffrement>",
    "sessionId": "<identifiant_de_la_session_de_chiffrement>",
    "createdAt": "2023-10-27T09:59:00Z"
  }
]
```

## Messages WebSocket

### Connexion

La connexion WebSocket est établie à l'URL suivante :
```
wss://api.hoomi.app/ws
```

L'authentification se fait en passant le token JWT dans le handshake :
```
Authorization: Bearer <JWT>
```

### Format des Messages WebSocket

Tous les messages échangés via WebSocket sont des objets JSON avec la structure suivante :
```json
{
  "type": "string", // Type du message
  "payload": {},    // Contenu du message
  "timestamp": "2023-10-27T10:00:00Z" // Horodatage
}
```

### Messages du Client vers le Serveur

#### Abonnement à un Hoomi
```json
{
  "type": "subscribe",
  "payload": {
    "hoomiId": "uuid-du-hoomi"
  },
  "timestamp": "2023-10-27T10:00:00Z"
}
```

#### Désabonnement d'un Hoomi
```json
{
  "type": "unsubscribe",
  "payload": {
    "hoomiId": "uuid-du-hoomi"
  },
  "timestamp": "2023-10-27T10:00:00Z"
}
```

#### Envoi d'un Message (via WebSocket)
```json
{
  "type": "sendMessage",
  "payload": {
    "hoomiId": "uuid-du-hoomi",
    "content": "<contenu_chiffré_en_base64>",
    "nonce": "<nonce_utilisé_pour_le_chiffrement>",
    "sessionId": "<identifiant_de_la_session_de_chiffrement>"
  },
  "timestamp": "2023-10-27T10:00:00Z"
}
```

### Messages du Serveur vers le Client

#### Confirmation d'Abonnement
```json
{
  "type": "subscribed",
  "payload": {
    "hoomiId": "uuid-du-hoomi",
    "status": "success"
  },
  "timestamp": "2023-10-27T10:00:00Z"
}
```

#### Nouveau Message Reçu
```json
{
  "type": "newMessage",
  "payload": {
    "message": {
      "id": "uuid-du-message",
      "hoomiId": "uuid-du-hoomi",
      "senderId": "uuid-de-l-expediteur",
      "senderName": "Nom de l'expéditeur",
      "content": "<contenu_chiffré_en_base64>",
      "nonce": "<nonce_utilisé_pour_le_chiffrement>",
      "sessionId": "<identifiant_de_la_session_de_chiffrement>",
      "createdAt": "2023-10-27T10:00:00Z"
    }
  },
  "timestamp": "2023-10-27T10:00:00Z"
}
```

#### Notification de Typing
```json
{
  "type": "typing",
  "payload": {
    "hoomiId": "uuid-du-hoomi",
    "userId": "uuid-de-l-utilisateur",
    "userName": "Nom de l'utilisateur",
    "isTyping": true
  },
  "timestamp": "2023-10-27T10:00:00Z"
}
```

#### Erreur
```json
{
  "type": "error",
  "payload": {
    "code": "string",
    "message": "Description de l'erreur"
  },
  "timestamp": "2023-10-27T10:00:00Z"
}
```

#### Message de Connexion/Déconnexion
```json
{
  "type": "presence",
  "payload": {
    "hoomiId": "uuid-du-hoomi",
    "userId": "uuid-de-l-utilisateur",
    "userName": "Nom de l'utilisateur",
    "status": "connected" // ou "disconnected"
  },
  "timestamp": "2023-10-27T10:00:00Z"
}
```

## Gestion des Erreurs

### Codes d'Erreur API

- `400 Bad Request` : Données du message invalides.
- `401 Unauthorized` : Token d'authentification invalide ou manquant.
- `403 Forbidden` : L'utilisateur n'est pas membre du Hoomi.
- `404 Not Found` : Hoomi non trouvé.
- `429 Too Many Requests` : Limite de taux dépassée.
- `500 Internal Server Error` : Erreur interne du serveur.

### Codes d'Erreur WebSocket

- `INVALID_TOKEN` : Le token d'authentification est invalide.
- `ALREADY_SUBSCRIBED` : Le client est déjà abonné à ce Hoomi.
- `NOT_SUBSCRIBED` : Le client n'est pas abonné à ce Hoomi.
- `INVALID_MESSAGE` : Le format du message est invalide.
- `NOT_A_MEMBER` : L'utilisateur n'est pas membre du Hoomi ciblé.

## Séquence Typique d'Interaction

1. **Connexion** : Le client établit une connexion WebSocket avec un token valide.
2. **Abonnement** : Le client envoie un message `subscribe` pour chaque Hoomi qu'il souhaite suivre.
3. **Réception d'historique** : Le client récupère l'historique via l'API REST.
4. **Envoi de message** : Le client peut envoyer un message via REST ou WebSocket.
5. **Réception en temps réel** : Le client reçoit les nouveaux messages via WebSocket.
6. **Indicateur de saisie** : Le client peut envoyer des messages `typing` pour indiquer qu'il est en train d'écrire.
7. **Déconnexion** : Le client se déconnecte de la WebSocket.

## Sécurité du Protocole

- Tous les échanges se font via une connexion sécurisée (WSS pour WebSocket, HTTPS pour REST).
- Les messages sont chiffrés de bout en bout avant d'être envoyés, donc le protocole transporte toujours du contenu chiffré.
- L'authentification est requise pour tous les points d'accès.
- Les tokens JWT ont une durée de vie limitée.