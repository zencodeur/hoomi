# Architecture du Système de Jeux Familiaux Hoomi

Je documente ici l'architecture que j'ai conçue pour le système de jeux familiaux d'Hoomi.

## Vue d'Ensemble

Le système de jeux de Hoomi est conçu pour être :
- **Simple d'utilisation** : Interface intuitive pour tous les âges.
- **Sécurisé** : Intégration complète avec le système d'authentification et d'autorisation.
- **Extensible** : Facile d'ajouter de nouveaux jeux.
- **Multijoueur** : Support des parties entre membres d'un même Hoomi.
- **Léger** : Ne pas surcharger l'application principale.

## Architecture Générale

### Structure en Couches

```
┌─────────────────────────────────────┐
│            UI / Interface           │
├─────────────────────────────────────┤
│      Logique des Jeux (Frontend)    │
├─────────────────────────────────────┤
│      API des Jeux (Backend)         │
├─────────────────────────────────────┤
│      Gestion des Parties            │
├─────────────────────────────────────┤
│        Stockage des Données         │
└─────────────────────────────────────┘
```

## Architecture Frontend (Web et Mobile)

### Approche

Les jeux sont implémentés comme des **applications web légères** qui s'intègrent dans l'interface principale (iframe ou composant webview).

### Technologies

- **Langages** : HTML5, CSS3, JavaScript (ES6+)
- **Frameworks** : Aucun framework lourd imposé, mais utilisation de bibliothèques légères si nécessaire
- **Communication** : API REST/WebSocket pour interagir avec le backend de jeu

### Intégration

- **Web** : Intégration via un conteneur HTML `<iframe>` ou un composant `<webview>`.
- **Mobile** : Utilisation d'un WebView natif (WKWebView sur iOS, WebView sur Android).

### Cycle de Vie d'un Jeu

1. **Chargement** : Le jeu est chargé dans un conteneur sécurisé.
2. **Initialisation** : Communication avec l'API de jeu pour récupérer l'état initial.
3. **Interaction** : L'utilisateur interagit avec le jeu.
4. **Synchronisation** : Les états du jeu sont synchronisés en temps réel entre les joueurs.
5. **Sauvegarde** : L'état du jeu est sauvegardé périodiquement.
6. **Fermeture** : Le jeu est arrêté et les ressources libérées.

## Architecture Backend

### Services

#### Service de Gestion des Jeux

- **Fonction** : Gérer le catalogue des jeux disponibles.
- **Endpoints** :
  - `GET /games` : Liste des jeux disponibles.
  - `GET /games/{id}` : Détails d'un jeu spécifique.

#### Service de Parties

- **Fonction** : Gérer les instances de parties en cours.
- **Endpoints** :
  - `POST /games/{id}/sessions` : Créer une nouvelle partie.
  - `GET /sessions/{id}` : Récupérer l'état d'une partie.
  - `POST /sessions/{id}/actions` : Envoyer une action dans la partie.
  - `DELETE /sessions/{id}` : Terminer une partie.

#### Service de Stockage des États

- **Fonction** : Persister les états des parties.
- **Stockage** : Base de données (PostgreSQL) avec une table dédiée `game_sessions`.

### Communication Temps Réel

- **WebSocket** : Utilisation de WebSocket pour la synchronisation en temps réel des états de jeu entre les joueurs.
- **Protocole** : Messages JSON avec des types prédéfinis (MOVE, UPDATE, CHAT, etc.).

## Structure d'un Jeu

### Fichiers Requis

```
game-name/
├── index.html          # Point d'entrée du jeu
├── css/
│   └── style.css       # Styles du jeu
├── js/
│   ├── game.js         # Logique principale du jeu
│   └── network.js      # Communication avec l'API
├── assets/
│   ├── images/         # Images du jeu
│   └── sounds/         # Sons du jeu (si applicable)
└── manifest.json       # Métadonnées du jeu
```

### Manifeste (manifest.json)

```json
{
  "id": "unique-game-id",
  "name": "Nom du Jeu",
  "version": "1.0.0",
  "description": "Description courte du jeu",
  "author": "Nom de l'auteur",
  "minPlayers": 2,
  "maxPlayers": 4,
  "category": "stratégie",
  "assets": {
    "icon": "assets/icon.png",
    "screenshots": ["assets/screenshot1.png"]
  }
}
```

## Gestion des Parties

### Création d'une Partie

1. Un membre du Hoomi sélectionne un jeu.
2. Il choisit les autres participants (membres du Hoomi).
3. Le backend crée une nouvelle session de jeu.
4. Un identifiant unique de session est généré.
5. Tous les participants reçoivent une notification.

### Rejoindre une Partie

1. Les participants cliquent sur la notification ou accèdent à la liste des parties en cours.
2. Le jeu est chargé avec l'identifiant de session.
3. Le joueur est ajouté à la partie.
4. L'état actuel du jeu est synchronisé.

### Déroulement d'une Partie

1. Les joueurs interagissent avec le jeu via l'interface.
2. Chaque action est envoyée à l'API de jeu.
3. L'API valide l'action et met à jour l'état de la partie.
4. Le nouvel état est diffusé en temps réel aux autres joueurs.
5. Le jeu se met à jour localement sur chaque appareil.

### Terminer une Partie

1. La partie se termine naturellement (victoire, égalité) ou manuellement.
2. L'état final est sauvegardé.
3. Les scores sont enregistrés.
4. Les participants sont notifiés de la fin de la partie.

## Sécurité

### Isolation des Jeux

- Chaque jeu s'exécute dans un environnement isolé (iframe/webview).
- Restrictions CORS strictes.
- Accès limité aux APIs du navigateur.

### Authentification

- Toutes les communications avec l'API de jeu nécessitent un token JWT valide.
- Le token est vérifié pour s'assurer que l'utilisateur est membre du Hoomi de la partie.

### Validation des Actions

- Le backend valide toutes les actions des joueurs.
- Prévention des triches basiques (mouvements invalides, etc.).

## Extensibilité

### Ajout de Nouveaux Jeux

1. Développement du jeu selon la structure définie.
2. Test du jeu en isolation.
3. Ajout du jeu au catalogue via l'API de gestion des jeux.
4. Le jeu devient disponible pour tous les utilisateurs.

### Mise à Jour des Jeux

1. Les développeurs peuvent publier de nouvelles versions.
2. Le système de versioning permet de gérer les mises à jour.
3. Les parties en cours continuent avec l'ancienne version jusqu'à leur terme.

## Performance

### Chargement

- Optimisation des assets (compression d'images, minification du code).
- Chargement progressif des ressources.

### Synchronisation

- Messages WebSocket légers.
- Mise à jour incrémentale de l'état du jeu.

### Stockage

- État des parties stocké de manière efficace.
- Nettoyage automatique des parties inactives après une période définie.

## Tests

### Tests Unitaires

- Pour la logique de jeu côté frontend.
- Pour les APIs de jeu côté backend.

### Tests d'Intégration

- Pour la communication entre frontend et backend.
- Pour la synchronisation multi-joueurs.

### Tests de Charge

- Pour s'assurer que le système peut gérer plusieurs parties simultanées.

Cette architecture modulaire et extensible permet d'ajouter facilement de nouveaux jeux tout en maintenant une expérience utilisateur fluide et sécurisée.