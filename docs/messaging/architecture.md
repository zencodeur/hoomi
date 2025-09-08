# Architecture du Système de Messagerie Hoomi

Je documente ici l'architecture du système de messagerie que j'ai conçue pour Hoomi.

## Vue d'Ensemble

Le système de messagerie de Hoomi est conçu pour être :
- **Sécurisé** : Avec un chiffrement de bout en bout intégré.
- **Temps réel** : Pour une expérience de conversation fluide.
- **Fiable** : Pour garantir la livraison des messages.
- **Évolutif** : Pour gérer un grand nombre d'utilisateurs et de conversations.

## Composants Principaux

### Service de Messagerie (Go)

C'est le cœur du système, responsable de :
- Recevoir les nouveaux messages des clients.
- Valider et traiter les messages.
- Coordonner le chiffrement E2EE.
- Diffuser les messages aux destinataires.
- Stocker les messages dans la base de données.

### WebSockets

J'utilise WebSockets pour maintenir des connexions persistantes entre le serveur et les clients, permettant la messagerie en temps réel.

### Base de Données (PostgreSQL)

Stockage persistant des messages, avec :
- La table `messages` pour les métadonnées.
- Le contenu chiffré est stocké dans le champ `content`.
- Index optimisés pour la récupération rapide des historiques.

### Broker de Messages (Optionnel)

Pour une architecture plus découplée, je pourrais introduire un broker de messages (comme Redis Pub/Sub ou Apache Kafka) pour gérer la distribution des messages entre les différents services.

## Flux de Données

### Envoi d'un Message

1. **Client** : L'utilisateur tape un message et clique sur "Envoyer".
2. **Client** : Le message est chiffré avec la clé de session E2EE.
3. **Client** : Le message chiffré est envoyé via HTTPS à l'API (`POST /hoomis/{id}/messages`).
4. **API** : Le service de messagerie reçoit la requête.
5. **API** : Le message est validé (authentification, appartenance au Hoomi).
6. **API** : Le message chiffré est stocké dans PostgreSQL.
7. **API** : Le message est diffusé en temps réel via WebSockets à tous les membres connectés du Hoomi.
8. **Clients** : Les clients reçoivent le message via WebSocket.
9. **Clients** : Les clients déchiffrent le message avec leur clé de session.

### Réception d'un Message (Temps Réel)

1. **Client** : Maintient une connexion WebSocket ouverte avec le serveur.
2. **Serveur** : Lorsqu'un nouveau message est envoyé à un Hoomi, il est poussé via WebSocket.
3. **Client** : Reçoit le message chiffré.
4. **Client** : Déchiffre le message avec la clé de session E2EE.
5. **Client** : Affiche le message dans l'interface.

### Récupération de l'Historique

1. **Client** : Lors du chargement d'un Hoomi, demande l'historique des messages (`GET /hoomis/{id}/messages`).
2. **API** : Récupère les messages chiffrés de la base de données.
3. **API** : Renvoie la liste des messages chiffrés.
4. **Client** : Déchiffre chaque message avec la clé de session.
5. **Client** : Affiche l'historique.

## Intégration E2EE

Le chiffrement de bout en bout est intégré à chaque étape :
- **Chiffrement côté client** : Avant l'envoi via l'API.
- **Stockage chiffré** : Le serveur ne voit jamais le contenu en clair.
- **Diffusion chiffrée** : Même les messages en temps réel sont chiffrés.
- **Déchiffrement côté client** : Uniquement sur les appareils des membres du Hoomi.

## Gestion des Connexions WebSocket

### Authentification

Les connexions WebSocket sont authentifiées via un token JWT passé lors de l'établissement de la connexion.

### Abonnement aux Hoomis

Une fois connecté, un client s'abonne aux Hoomis qui l'intéressent en envoyant un message spécifique via la WebSocket.

### Gestion des Déconnexions

Le serveur gère les déconnexions (volontaires ou involontaires) et met à jour la liste des membres connectés en conséquence.

## Scalabilité

### Horizontal Scaling

Le service de messagerie peut être déployé sur plusieurs instances. Un load balancer distribue les connexions WebSocket et les requêtes API.

### Limitations Actuelles

- Toutes les connexions WebSocket d'un Hoomi doivent être gérées par la même instance (ou un système de partage d'état comme Redis).
- Pour une très grande scalabilité, un broker de messages dédié serait nécessaire.

## Tolérance aux Pannes

### Persistance

Les messages sont stockés de manière persistante dans PostgreSQL, garantissant qu'ils ne sont pas perdus même si le serveur redémarre.

### Redondance

Le système est conçu pour fonctionner avec des déploiements redondants (plusieurs instances du service de messagerie).

## Surveillance

Je prévois d'implémenter des métriques pour surveiller :
- Le nombre de messages envoyés/reçus par seconde.
- La latence de livraison des messages.
- Le nombre de connexions WebSocket actives.
- Les taux d'erreur.